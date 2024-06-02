package server

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	lookup "github.com/brotherlogic/mdb/lookup"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	ghbclient "github.com/brotherlogic/githubridge/client"
	rsclient "github.com/brotherlogic/rstore/client"

	ghbpb "github.com/brotherlogic/githubridge/proto"
	pb "github.com/brotherlogic/mdb/proto"
	rspb "github.com/brotherlogic/rstore/proto"
)

const (
	MDB_PATH     = "github.com/brotherlogic/mdb"
	GHB_PASSWORD = "ghbridge_password"

	REFILL_FREQUENCY = time.Minute * 5
)

var (
	validationError = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "mdb_validate_error",
	}, []string{"error"})
	refillError = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "mdb_refill_error",
	}, []string{"error"})
	types = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mdb_machine_types",
	}, []string{"type"})
)

func metrics(config *pb.Mdb) {
	typeCount := make(map[string]float64)
	for _, machine := range config.GetMachines() {
		typeCount[fmt.Sprintf("%v", machine.GetType())]++
	}
	for key, val := range typeCount {
		types.With(prometheus.Labels{"type": key}).Set(val)
	}
}

type Server struct {
	ghbclient ghbclient.GithubridgeClient
	rsclient  rsclient.RStoreClient
	running   bool
}

func NewServer(ctx context.Context) *Server {
	ghbclient, err := ghbclient.GetClientInternal()
	if err != nil {
		log.Fatalf("Bad client get: %v", err)
	}
	rsclient, err := rsclient.GetClient()
	if err != nil {
		log.Fatalf("Unable to get rstore client: %v", err)
	}
	return &Server{ghbclient: ghbclient, rsclient: rsclient, running: true}
}

func (s *Server) RunRefillLoop() {
	for s.running {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		defer cancel()

		err := s.refillDatabase(ctx)
		if err != nil {
			refillError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)})
			log.Fatalf("error building machine database on init: %v", err)
		}

		time.Sleep(REFILL_FREQUENCY)
	}
}

func (s *Server) refillDatabase(ctx context.Context) error {
	config, err := s.loadConfig(ctx)
	if err != nil {
		return fmt.Errorf("unable to load config: %w", err)
	}

	metrics(config)

	err = lookup.FillDB(ctx, config)
	if err != nil {
		return fmt.Errorf("unable to fill the db: %w", err)
	}

	// Validate machines if we need to
	if config.GetConfig().GetIssueId() == 0 {
		s.validateMachines(ctx, config)
	} else {
		s.checkIssue(ctx, config)
	}

	return s.saveConfig(ctx, config)
}

func cleanConfig(config *pb.Mdb) {
	for _, machine := range config.GetMachines() {
		if strings.Contains(machine.GetMac(), "Unable") {
			machine.Mac = ""
		}
	}

	var vm []*pb.Machine
	for _, machine := range config.GetMachines() {
		if machine.Ipv4 > 0 {
			vm = append(vm, machine)
		}
	}
	config.Machines = vm

	// Clear repeated entries
	for i, machine1 := range config.GetMachines() {
		for j, machine2 := range config.GetMachines() {
			if i != j {
				if machine1.GetIpv4() == machine2.GetIpv4() && machine1.GetHostname() == machine2.GetHostname() {
					if machine1.GetMac() != "" && machine2.GetMac() == "" {
						machine2.MarkedForDelete = true
					} else if machine2.GetMac() != "" && machine1.GetMac() == "" {
						machine1.MarkedForDelete = true
					}
				}
			}
		}
	}

	var nm []*pb.Machine
	for _, machine := range config.GetMachines() {
		if !machine.GetMarkedForDelete() {
			nm = append(nm, machine)
		} else {
			log.Printf("Deleting %v", machine)
		}
	}
	config.Machines = nm
}

func (s *Server) loadConfig(ctx context.Context) (*pb.Mdb, error) {
	data, err := s.rsclient.Read(ctx, &rspb.ReadRequest{
		Key: MDB_PATH,
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return &pb.Mdb{}, nil
		}
		return nil, err
	}

	ret := &pb.Mdb{}
	err = proto.Unmarshal(data.GetValue().GetValue(), ret)
	cleanConfig(ret)
	return ret, err
}

func (s *Server) saveConfig(ctx context.Context, mdb *pb.Mdb) error {
	data, err := proto.Marshal(mdb)
	if err != nil {
		return err
	}
	_, err = s.rsclient.Write(ctx, &rspb.WriteRequest{
		Key:   MDB_PATH,
		Value: &anypb.Any{Value: data},
	})
	return err
}

func (s *Server) ListMachines(ctx context.Context, req *pb.ListMachinesRequest) (*pb.ListMachinesResponse, error) {
	config, err := s.loadConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListMachinesResponse{Machines: config.GetMachines()}, nil
}

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func (s *Server) raiseIssue(ctx context.Context, mdb *pb.Mdb, machine *pb.Machine, verr pb.MachineErrors) error {
	body := ""

	ccount := 0
	for _, m := range mdb.GetMachines() {
		if m.GetController() == machine.GetController() {
			ccount++
		}
	}

	switch verr {
	case pb.MachineErrors_MACHINE_ERROR_MISSING_TYPE:
		body = fmt.Sprintf("%v (%v) is missing the machine type [%v]", machine.GetHostname(), machine.GetController(), ccount)
	case pb.MachineErrors_MACHINE_ERROR_MISSING_USE:
		body = fmt.Sprintf("%v (%v) is missing the machine use", machine.GetHostname(), machine.GetType())
	case pb.MachineErrors_MACHINE_ERROR_UNSTABLE_IP:
		body = fmt.Sprintf("%v has recorded mulitple IPs (e.g. %v)", machine.GetHostname(), ipv4ToString(machine.GetIpv4()))
	case pb.MachineErrors_MACHINE_ERROR_CLUSTER_MISSING_TYPE:
		body = fmt.Sprintf("%v is missing the cluster type", machine.GetHostname())
	case pb.MachineErrors_MACHINE_ERROR_NONE:
		return status.Errorf(codes.Internal, "Trying to raise issue for unbroken machine")
	}

	issue, err := s.ghbclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{
		User:  "brotherlogic",
		Repo:  "mdb",
		Title: "Missing data in MDB",
		Body:  body,
	})
	if err != nil {
		// If CreateIssue returns AlreadyExists we can re-use the issue id
		if status.Code(err) != codes.AlreadyExists || issue != nil {
			log.Printf("Returning %v maybe because %v", err, issue)
			return err
		}
	}

	log.Printf("Resolved to %v (%v)", issue, err)

	if mdb.GetConfig() == nil {
		mdb.Config = &pb.Config{}
	}

	mdb.GetConfig().CurrentMachine = machine
	mdb.GetConfig().IssueId = int32(issue.GetIssueId())
	mdb.GetConfig().IssueType = verr

	return nil
}

func (s *Server) dataMissing(ctx context.Context, machine *pb.Machine) pb.MachineErrors {
	if machine.GetType() == pb.MachineType_MACHINE_TYPE_UNKNOWN {
		return pb.MachineErrors_MACHINE_ERROR_MISSING_TYPE
	}

	if machine.GetType() == pb.MachineType_MACHINE_TYPE_INTEL ||
		machine.GetType() == pb.MachineType_MACHINE_TYPE_RASPBERRY_PI ||
		machine.GetType() == pb.MachineType_MACHINE_TYPE_AMD {
		if machine.GetUse() == pb.MachineUse_MACHINE_USE_UNKNOWN {
			return pb.MachineErrors_MACHINE_ERROR_MISSING_USE
		}
	}

	if machine.GetUse() == pb.MachineUse_MACHINE_USE_KUBERNETES_CLUSTER {
		if machine.GetClusterType() == pb.ClusterType_CLUSTER_TYPE_UNKNONW {
			return pb.MachineErrors_MACHINE_ERROR_CLUSTER_MISSING_TYPE
		}
	}

	return pb.MachineErrors_MACHINE_ERROR_NONE
}

func resolveController(controller string) pb.MachineType {
	switch controller {
	case "(Raspberry Pi Trading)":
		return pb.MachineType_MACHINE_TYPE_RASPBERRY_PI
	case "(Raspberry Pi Foundation)":
		return pb.MachineType_MACHINE_TYPE_RASPBERRY_PI
	case "(Belkin International)":
		return pb.MachineType_MACHINE_TYPE_IOT_DEVICE
	case "(Sonos)":
		return pb.MachineType_MACHINE_TYPE_IOT_DEVICE
	case "(Miele & Cie. KG)":
		return pb.MachineType_MACHINE_TYPE_IOT_DEVICE
	}

	return pb.MachineType_MACHINE_TYPE_UNKNOWN
}

func (s *Server) autofill(ctx context.Context, machine *pb.Machine) {
	if machine.GetType() == pb.MachineType_MACHINE_TYPE_UNKNOWN {
		machine.Type = resolveController(machine.GetController())
	}
}

func (s *Server) validateMachines(ctx context.Context, mdb *pb.Mdb) error {
	log.Printf("Validating Machines")
	for _, machine := range mdb.GetMachines() {
		s.autofill(ctx, machine)

		valid := s.dataMissing(ctx, machine)
		if valid != pb.MachineErrors_MACHINE_ERROR_NONE {
			err := s.raiseIssue(ctx, mdb, machine, valid)
			log.Printf("Found issue with %v -> %v with %v", machine, valid, err)
			validationError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)})
			return err
		}
	}

	log.Printf("Looking for repeated IP addresses")
	macToIP := make(map[string]int32)
	for _, machine := range mdb.GetMachines() {
		if _, ok := macToIP[machine.GetMac()]; ok {
			err := s.raiseIssue(ctx, mdb, machine, pb.MachineErrors_MACHINE_ERROR_UNSTABLE_IP)
			validationError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)})
			return err
		}
	}

	log.Printf("Validation Complete")
	return nil
}

func (s *Server) checkIssue(ctx context.Context, mdb *pb.Mdb) error {
	labels, err := s.ghbclient.GetLabels(ctx, &ghbpb.GetLabelsRequest{
		User: "brotherlogic",
		Repo: "mdb",
		Id:   mdb.GetConfig().GetIssueId(),
	})
	log.Printf("Checking issues: %v -> %v", mdb.GetConfig(), err)
	if err != nil {
		return err
	}

	for _, label := range labels.GetLabels() {
		switch label {
		case "raspberrypi":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_RASPBERRY_PI
		case "intel":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_INTEL
		case "iot":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_IOT_DEVICE
		case "apple":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_APPLE
		case "phone":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_PHONE
		case "tablet":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_TABLET
		case "amd":
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_AMD
		case "kubernetes-cluster":
			mdb.GetConfig().GetCurrentMachine().Use = pb.MachineUse_MACHINE_USE_KUBERNETES_CLUSTER
		case "development-server":
			mdb.GetConfig().GetCurrentMachine().Use = pb.MachineUse_MACHINE_USE_DEV_SERVER
		case "development-desktop":
			mdb.GetConfig().GetCurrentMachine().Use = pb.MachineUse_MACHINE_USE_DEV_DESKTOP
		case "home-cluster":
			mdb.GetConfig().GetCurrentMachine().Use = pb.MachineUse_MACHINE_USE_LOCAL_CLUSTER
		case "windows":
			mdb.GetConfig().GetCurrentMachine().Use = pb.MachineUse_MACHINE_USE_NOT_IN_USE
		case "pi-server":
			mdb.GetConfig().GetCurrentMachine().Use = pb.MachineUse_MACHINE_USE_PI_SERVER
		case "lead":
			mdb.GetConfig().GetCurrentMachine().ClusterType = pb.ClusterType_CLUSTER_TYPE_LEAD
		case "follow":
			mdb.GetConfig().GetCurrentMachine().ClusterType = pb.ClusterType_CLUSTER_TYPE_FOLLOWER
		case "fixed":
			// Clear all instances of this entity and re-create the db
			var nm []*pb.Machine
			for _, machine := range mdb.GetMachines() {
				if machine.GetMac() != mdb.GetConfig().GetCurrentMachine().GetMac() {
					nm = append(nm, machine)
				}
			}
			mdb.Machines = nm
		default:
			log.Printf("Skipping label %v on %v", label, mdb.GetConfig().GetCurrentMachine())
		}
	}

	//Resolve the machine
	return s.resolveMachine(ctx, mdb)
}

func (s *Server) resolveMachine(ctx context.Context, mdb *pb.Mdb) error {
	log.Printf("Resolving Machines: %v", mdb.GetConfig())
	if mdb.GetConfig().GetCurrentMachine().GetType() == pb.MachineType_MACHINE_TYPE_UNKNOWN {
		return nil
	}

	for _, machine := range mdb.GetMachines() {
		if machine.GetController() == mdb.GetConfig().GetCurrentMachine().GetController() && machine.GetHostname() == mdb.GetConfig().GetCurrentMachine().GetHostname() {

			if machine.GetType() == pb.MachineType_MACHINE_TYPE_UNKNOWN && mdb.GetConfig().GetCurrentMachine().GetType() != pb.MachineType_MACHINE_TYPE_UNKNOWN {
				machine.Type = mdb.GetConfig().GetCurrentMachine().GetType()
			}

			log.Printf("USE: %v, %v", machine.GetUse(), mdb.GetConfig().GetCurrentMachine().GetUse())
			if machine.GetUse() == pb.MachineUse_MACHINE_USE_UNKNOWN && mdb.GetConfig().GetCurrentMachine().GetUse() != pb.MachineUse_MACHINE_USE_UNKNOWN {
				machine.Use = mdb.GetConfig().GetCurrentMachine().GetUse()
			}

			if machine.GetClusterType() == pb.ClusterType_CLUSTER_TYPE_UNKNONW && mdb.GetConfig().GetCurrentMachine().GetClusterType() != pb.ClusterType_CLUSTER_TYPE_UNKNONW {
				machine.ClusterType = mdb.Config.CurrentMachine.GetClusterType()
			}

			if (mdb.GetConfig().GetIssueType() == pb.MachineErrors_MACHINE_ERROR_MISSING_TYPE && machine.GetType() != pb.MachineType_MACHINE_TYPE_UNKNOWN) ||
				(mdb.GetConfig().GetIssueType() == pb.MachineErrors_MACHINE_ERROR_MISSING_USE && machine.GetUse() != pb.MachineUse_MACHINE_USE_UNKNOWN) ||
				(mdb.GetConfig().GetIssueType() == pb.MachineErrors_MACHINE_ERROR_CLUSTER_MISSING_TYPE && machine.GetClusterType() != pb.ClusterType_CLUSTER_TYPE_UNKNONW) {
				_, err := s.ghbclient.CloseIssue(ctx, &ghbpb.CloseIssueRequest{
					User: "brotherlogic",
					Repo: "mdb",
					Id:   int64(mdb.GetConfig().GetIssueId()),
				})
				log.Printf("Resolved machine and closed issue: %v", err)
				if err == nil {
					mdb.GetConfig().CurrentMachine = nil
					mdb.GetConfig().IssueId = 0
					mdb.GetConfig().IssueType = pb.MachineErrors_MACHINE_ERROR_NONE
				}

				return err
			} else {
				// Issue was not resolved
				return nil
			}
		}
	}

	// If we can't find the machine, delete the issue
	log.Printf("Unable to locate machine in MDB, closing issue: %v", mdb.GetConfig().GetCurrentMachine())
	_, err := s.ghbclient.CloseIssue(ctx, &ghbpb.CloseIssueRequest{
		User: "brotherlogic",
		Repo: "mdb",
		Id:   int64(mdb.GetConfig().GetIssueId()),
	})
	if err == nil {
		mdb.GetConfig().CurrentMachine = nil
		mdb.GetConfig().IssueId = 0
	}

	return err
}

func (s *Server) UpdateMachine(ctx context.Context, req *pb.UpdateMachineRequest) (*pb.UpdateMachineResponse, error) {
	config, err := s.loadConfig(ctx)
	if err != nil {
		return nil, err
	}

	if req.GetIpv4() > 0 && req.GetRemove() {
		var machines []*pb.Machine
		for _, machine := range config.GetMachines() {
			if machine.GetIpv4() != req.GetIpv4() {
				machines = append(machines, machine)
			}
		}
		config.Machines = machines
		return &pb.UpdateMachineResponse{}, s.saveConfig(ctx, config)
	}

	for _, machine := range config.GetMachines() {

		if machine.GetHostname() == req.GetHostname() {
			updated := false
			if req.GetNewType() != pb.MachineType_MACHINE_TYPE_UNKNOWN {
				machine.Type = req.GetNewType()
				updated = true
			}
			if req.GetNewUse() != pb.MachineUse_MACHINE_USE_UNKNOWN {
				machine.Use = req.GetNewUse()
				update = true
			}

			if req.GetMarkUpdate() {
				machine.LastUpdated = time.Now().UnixNano()
				updated = true
			}

			if req.GetVersion() != "" {
				machine.Version = req.GetVersion()
				updated = true
			}

			if updated {
				return &pb.UpdateMachineResponse{}, s.saveConfig(ctx, config)
			}

			return nil, status.Errorf(codes.FailedPrecondition, "No update was specified: %v", req)
		}
	}

	return nil, status.Errorf(codes.NotFound, "machine %v was not found", req.GetHostname())
}
