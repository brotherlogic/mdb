package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/brotherlogic/mdb/server"

	pb "github.com/brotherlogic/mdb/proto"
)

var (
	port        = flag.Int("port", 8080, "Server port for grpc traffic")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	s := server.NewServer(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("mdb is unable to listen on the grpc port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterMDBServiceServer(gs, s)

	go func() {
		err := s.FillDB()
		if err != nil {
			log.Fatalf("error building machine database on init: %v", err)
		}
	}()

	err = gs.Serve(lis)
	log.Printf("mdb is unable to serve http: %v", err)
}
