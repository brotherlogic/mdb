package server

import (
	"encoding/binary"
	"fmt"
	"net"

	pb "github.com/brotherlogic/mdb/proto"
)

func (s *Server) fillDB() {
	toru := "192.168.86.22"
}

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func (s *Server) lookupv4(ipv4 uint32) (*pb.Machine, error) {
	addr, err := net.LookupAddr(ipv4ToString(ipv4))
	if err != nil {
		return nil, fmt.Errorf("lookupaddr error: %v", err)
	}

	return &pb.Machine{
		Ipv4:     ipv4,
		Hostname: addr[0],
	}, nil
}
