package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/brotherlogic/mdb/proto"
)

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func strToIpv4(ipv4 string) uint32 {
	return binary.BigEndian.Uint32(net.ParseIP(ipv4)[12:16])
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewMDBServiceClient(conn)

	if os.Args[2] == "delete" {
		resp, err := client.UpdateMachine(ctx, &pb.UpdateMachineRequest{Ipv4: strToIpv4(os.Args[3]), Remove: true})
		if err != nil {
			log.Printf("Delete: %v -> %v", resp, err)
		}
		return
	} else if os.Args[2] == "set_home_cluster" {
		_, err := client.UpdateMachine(ctx, &pb.UpdateMachineRequest{Ipv4: strToIpv4(os.Args[3]), NewUse: pb.MachineUse_MACHINE_USE_LOCAL_CLUSTER})
		if err != nil {
			log.Print("Home: %v", err)
			return
		}
	}

	resp, err := client.ListMachines(ctx, &pb.ListMachinesRequest{})
	if err != nil {
		log.Fatalf("Unable to drain queue: %v", err)
	}

	for _, machine := range resp.GetMachines() {
		fmt.Printf("%v %v\n", ipv4ToString(machine.GetIpv4()), machine)
	}
}
