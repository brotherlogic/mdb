package server

import (
	"context"

	pb "github.com/brotherlogic/mdb/proto"
)

type Server struct {
	machines []*pb.Machine
}

func (s *Server) ListMachines(ctx context.Context, req *pb.ListMachinesRequest) (*pb.ListMachinesResponse, error) {
	return &pb.ListMachinesResponse{Machines: s.machines}, nil
}

func NewServer(ctx context.Context) *Server {
	return &Server{}
}
