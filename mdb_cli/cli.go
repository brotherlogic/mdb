package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/brotherlogic/mdb/proto"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewMDBServiceClient(conn)

	resp, err := client.ListMachines(ctx, &pb.ListMachinesRequest{})
	if err != nil {
		log.Fatalf("Unable to drain queue: %v", err)
	}

	for _, machine := range resp.GetMachines() {
		fmt.Printf("%v\n", machine)
	}
}
