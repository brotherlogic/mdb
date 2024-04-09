package server

import (
	"context"
	"fmt"
	"log"
	"os"

	lookup "github.com/brotherlogic/mdb/lookup"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
)

var (
	validationError = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "mdb_validate_error",
	}, []string{"error"})
)

type Server struct {
	ghbclient ghbclient.GithubridgeClient
	rsclient  rsclient.RStoreClient
}

func (s *Server) RefillDatabase(ctx context.Context) error {
	config, err := s.loadConfig(ctx)
	if err != nil {
		return err
	}

	err = lookup.FillDB(ctx, config)
	if err != nil {
		return err
	}
	return s.saveConfig(ctx, config)
}

func (s *Server) loadConfig(ctx context.Context) (*pb.Mdb, error) {
	data, err := s.rsclient.Read(ctx, &rspb.ReadRequest{
		Key: MDB_PATH,
	})
	if err != nil {
		return nil, err
	}

	ret := &pb.Mdb{}
	err = proto.Unmarshal(data.GetValue().GetValue(), ret)
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

func NewServer(ctx context.Context) *Server {
	ghbclient, err := ghbclient.GetClientInternal(os.Getenv(GHB_PASSWORD))
	if err != nil {
		log.Fatalf("Bad client get: %v", err)
	}
	rsclient, err := rsclient.GetClient()
	if err != nil {
		log.Fatalf("Unable to get rstore client: %v", err)
	}
	return &Server{ghbclient: ghbclient, rsclient: rsclient}
}

func (s *Server) validateMachine(ctx context.Context, mdb *pb.Mdb, machine *pb.Machine) error {
	issue, err := s.ghbclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{
		User:  "brotherlogic",
		Repo:  "mdb",
		Title: "Missing data in MDB",
		Body:  fmt.Sprintf("%v is missing data", machine),
	})
	if err != nil {
		return err
	}
	mdb.GetConfig().CurrentMachine = machine
	mdb.GetConfig().IssueId = int32(issue.GetIssueId())

	return nil
}

func (s *Server) validateMachines(ctx context.Context, mdb *pb.Mdb) error {
	for _, machine := range mdb.GetMachines() {
		if machine.GetType() == pb.MachineType_MACHINE_TYPE_UNKNOWN {
			err := s.validateMachine(ctx, mdb, machine)
			validationError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)})
			return err
		}
	}

	return nil
}

func (s *Server) checkIssue(ctx context.Context, mdb *pb.Mdb) error {
	labels, err := s.ghbclient.GetLabels(ctx, &ghbpb.GetLabelsRequest{
		User: "brotherlogic",
		Repo: "mdb",
		Id:   mdb.GetConfig().GetIssueId(),
	})
	if err != nil {
		return err
	}

	for _, label := range labels.GetLabels() {
		if label == "raspberrypi" {
			mdb.GetConfig().GetCurrentMachine().Type = pb.MachineType_MACHINE_TYPE_RASPBERRY_PI
		}
	}

	//Resolve the machine
	return s.resolveMachine(ctx, mdb)
}

func (s *Server) resolveMachine(ctx context.Context, mdb *pb.Mdb) error {
	if mdb.GetConfig().GetCurrentMachine().GetType() == pb.MachineType_MACHINE_TYPE_UNKNOWN {
		return nil
	}

	for _, machine := range mdb.GetMachines() {
		if machine.GetController() == mdb.GetConfig().GetCurrentMachine().GetController() && machine.GetHostname() == mdb.GetConfig().GetCurrentMachine().GetHostname() {
			machine.Type = mdb.GetConfig().GetCurrentMachine().GetType()
			return nil
		}
	}

	mdb.Machines = append(mdb.Machines, mdb.GetConfig().GetCurrentMachine())
	return nil
}
