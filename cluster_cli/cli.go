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

	resp, err := client.ListMachines(ctx, &pb.ListMachinesRequest{})
	if err != nil {
		log.Fatalf("Unable to drain queue: %v", err)
	}

	fmt.Println("[master]")
	for _, machine := range resp.GetMachines() {
		if machine.GetUse() == pb.MachineUse_MACHINE_USE_KUBERNETES_CLUSTER && machine.GetClusterType() == pb.ClusterType_CLUSTER_TYPE_LEAD {
			fmt.Printf("%v # %v\n", ipv4ToString(machine.GetIpv4()), machine.GetHostname())
			break
		}
	}

	fmt.Printf("\n[node]\n")
	for _, machine := range resp.GetMachines() {
		if machine.GetUse() == pb.MachineUse_MACHINE_USE_KUBERNETES_CLUSTER && machine.GetClusterType() == pb.ClusterType_CLUSTER_TYPE_FOLLOWER {
			fmt.Printf("%v # %v\n", ipv4ToString(machine.GetIpv4()), machine.GetHostname())
		}
	}

	fmt.Printf("\n[k3s_cluster:children]\nmaster\nnode")
}
