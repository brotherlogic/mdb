package server

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"

	ghbpb "github.com/brotherlogic/githubridge/proto"
	pb "github.com/brotherlogic/mdb/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	LOWER = 1
	UPPER = 256

	machinesFound = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "mdb_machine_count",
	})

	lookupError = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "mdb_lookup_error",
	}, []string{"error"})

	validationError = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "mdb_validate_error",
	}, []string{"error"})
)

func (s *Server) getMacAddress(addr string) (string, string) {
	out, err := exec.Command("/usr/bin/nmap", addr).CombinedOutput()
	if err != nil {
		return fmt.Sprintf("unable to nmap: %v -> %v", err, string(out)), ""
	}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "MAC") {
			return strings.Split(line, " ")[2], strings.Join(strings.Split(line, " ")[3:], " ")
		}
	}

	return fmt.Sprintf("Unable to find mac address: %v", string(out)), ""
}

func mergeInto(machines []*pb.Machine, machine *pb.Machine) []*pb.Machine {
	for _, exMachine := range machines {
		if exMachine.Controller == machine.Controller && exMachine.Hostname == machine.Hostname {
			exMachine.Ipv4 = machine.Ipv4
			exMachine.Mac = machine.Mac
			return machines
		}
	}

	return append(machines, machine)
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

func (s *Server) FillDB(ctx context.Context) error {
	mdb, err := s.loadConfig(ctx)
	if err != nil {
		return err
	}

	// Don't run a fill if we're working on a machine
	if mdb.GetConfig().GetIssueId() != 0 {
		return s.checkIssue(ctx, mdb)
	}

	for i := LOWER; i <= UPPER; i++ {
		ipv4 := fmt.Sprintf("192.168.86.%v", i)
		machine, err := s.lookupv4str(ipv4)
		if err != nil {
			lookupError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)}).Inc()
		} else {
			mac, thing := s.getMacAddress(ipv4)
			log.Printf("%v -> %v, %v", ipv4, mac, thing)
			machine.Mac = mac
			machine.Controller = thing
			mdb.Machines = mergeInto(mdb.Machines, machine)
		}
	}

	log.Printf("Setting machine count: %v", len(mdb.Machines))
	machinesFound.Set(float64(len(mdb.Machines)))

	s.validateMachines(ctx, mdb)

	return nil
}

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func strToIpv4(ipv4 string) uint32 {
	return binary.BigEndian.Uint32(net.ParseIP(ipv4))
}

func (s *Server) lookupv4(ipv4 uint32) (*pb.Machine, error) {
	return s.lookupv4str(ipv4ToString(ipv4))
}

func (s *Server) lookupv4str(ipv4 string) (*pb.Machine, error) {
	addr, err := net.LookupAddr(ipv4)
	if err != nil {
		return nil, fmt.Errorf("lookupaddr error: (%v) -> %v", ipv4, err)
	}

	log.Printf("FOUND %v -> %v", ipv4, addr)

	return &pb.Machine{
		Ipv4:     strToIpv4(ipv4),
		Hostname: addr[0],
	}, nil
}
