package server

import (
	"context"
	"testing"

	ghbclient "github.com/brotherlogic/githubridge/client"
	ghbpb "github.com/brotherlogic/githubridge/proto"
	rsclient "github.com/brotherlogic/pstore/client"

	pb "github.com/brotherlogic/mdb/proto"
)

func GetTestServer(machines []*pb.Machine) (*Server, ghbclient.GithubridgeClient) {
	ghc := ghbclient.GetTestClient()
	s := &Server{
		ghbclient: ghc,
		psclient:  rsclient.GetTestClient(),
	}

	err := s.saveConfig(context.Background(), &pb.Mdb{Machines: machines})
	if err != nil {
		panic(err)
	}

	return s, ghc
}

func TestSetCTypeClearsIssue(t *testing.T) {
	s, ghc := GetTestServer([]*pb.Machine{
		{
			Ipv4:     1234,
			Hostname: "blah",
			Mac:      "MAC",
			Type:     pb.MachineType_MACHINE_TYPE_AMD,
			Use:      pb.MachineUse_MACHINE_USE_KUBERNETES_CLUSTER,
		},
	})

	// Validate the mdb
	mdb, err := s.loadConfig(context.Background())
	if err != nil {
		t.Fatalf("Unablet to load config: %v", err)
	}
	s.validateMachines(context.Background(), mdb)

	if mdb.GetConfig().GetIssueId() == 0 {
		t.Errorf("Issue was not created for missing cluster type: %v", mdb.GetConfig())
	}
	s.saveConfig(context.Background(), mdb)

	// Create the label
	ghc.AddLabel(context.Background(), &ghbpb.AddLabelRequest{
		User:  "brotherlogic",
		Repo:  "mdb",
		Id:    mdb.GetConfig().GetIssueId(),
		Label: "lead",
	})

	s.checkIssue(context.Background(), mdb)

	if mdb.GetConfig().GetIssueId() != 0 {
		t.Errorf("Issue was not removed on label set")
	}
}

func TestSetTypeClearsIssue(t *testing.T) {
	s, ghc := GetTestServer([]*pb.Machine{
		{
			Ipv4:     1234,
			Hostname: "blah",
			Mac:      "MAC",
		},
	})

	// Validate the mdb
	mdb, err := s.loadConfig(context.Background())
	if err != nil {
		t.Fatalf("Unablet to load config: %v", err)
	}
	s.validateMachines(context.Background(), mdb)

	if mdb.GetConfig().GetIssueId() == 0 {
		t.Errorf("Issue was not created for missing type: %v", mdb.GetConfig())
	}
	s.saveConfig(context.Background(), mdb)

	// Create the label
	ghc.AddLabel(context.Background(), &ghbpb.AddLabelRequest{
		User:  "brotherlogic",
		Repo:  "mdb",
		Id:    mdb.GetConfig().GetIssueId(),
		Label: "iot",
	})

	s.checkIssue(context.Background(), mdb)

	if mdb.GetConfig().GetIssueId() != 0 {
		t.Errorf("Issue was not removed on label set")
	}
}

func TestSetUseClearsIssue(t *testing.T) {
	s, ghc := GetTestServer([]*pb.Machine{
		{
			Ipv4:     1234,
			Hostname: "blah",
			Mac:      "MAC",
			Type:     pb.MachineType_MACHINE_TYPE_RASPBERRY_PI,
		},
	})

	// Validate the mdb
	mdb, err := s.loadConfig(context.Background())
	if err != nil {
		t.Fatalf("Unablet to load config: %v", err)
	}
	s.validateMachines(context.Background(), mdb)

	if mdb.GetConfig().GetIssueId() == 0 {
		t.Errorf("Issue was not created for missing use: %v", mdb.GetConfig())
	}
	s.saveConfig(context.Background(), mdb)

	// Create the label
	ghc.AddLabel(context.Background(), &ghbpb.AddLabelRequest{
		User:  "brotherlogic",
		Repo:  "mdb",
		Id:    mdb.GetConfig().GetIssueId(),
		Label: "kubernetes-cluster",
	})

	s.checkIssue(context.Background(), mdb)

	if mdb.GetConfig().GetIssueId() != 0 {
		t.Errorf("Issue was not removed on label set")
	}
}

func TestCleanMachines(t *testing.T) {
	s, _ := GetTestServer([]*pb.Machine{
		{
			Ipv4:     1234,
			Hostname: "blah",
			Mac:      "MAC",
			Type:     pb.MachineType_MACHINE_TYPE_RASPBERRY_PI,
		},
		{
			Ipv4:     1234,
			Hostname: "blah",
		},
	})

	// Validate the mdb
	mdb, err := s.loadConfig(context.Background())
	if err != nil {
		t.Fatalf("Unablet to load config: %v", err)
	}

	if len(mdb.GetMachines()) == 2 {
		t.Errorf("Machine was not cleaned")
	}

}

func TestCleanMachines_WithDiffHostname(t *testing.T) {
	s, _ := GetTestServer([]*pb.Machine{
		{
			Ipv4:     1234,
			Hostname: "blah",
			Mac:      "MAC",
			Type:     pb.MachineType_MACHINE_TYPE_RASPBERRY_PI,
		},
		{
			Ipv4:     1234,
			Hostname: "foo",
		},
	})

	// Validate the mdb
	mdb, err := s.loadConfig(context.Background())
	if err != nil {
		t.Fatalf("Unablet to load config: %v", err)
	}

	if len(mdb.GetMachines()) != 2 {
		t.Errorf("Machine was cleaned")
	}

}
