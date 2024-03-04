package server

import "context"

type Server struct{}

func NewServer(ctx context.Context) *Server {
	return &Server{}
}
