package server

import (
	"context"

	pb "github.com/brotherlogic/mdb/proto"
)

type Server struct {
	machines []*pb.Machine
}

func NewServer(ctx context.Context) *Server {
	return &Server{}
}
