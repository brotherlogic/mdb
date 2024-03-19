package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	ghbclient "github.com/brotherlogic/githubridge/client"
	rsclient "github.com/brotherlogic/rstore/client"

	pb "github.com/brotherlogic/mdb/proto"
	rspb "github.com/brotherlogic/rstore/proto"
)

const (
	MDB_PATH = "github.com/brotherlogic/mdb"
)

type Server struct {
	ghbclient ghbclient.GithubridgeClient
	rsclient  rsclient.RStoreClient
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
	ghbclient, err := ghbclient.GetClientInternal()
	if err != nil {
		log.Fatalf("Bad client get: %v", err)
	}
	rsclient, err := rsclient.GetClient()
	if err != nil {
		log.Fatalf("Unable to get rstore client: %v", err)
	}
	return &Server{ghbclient: ghbclient, rsclient: rsclient}
}
