package server

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	pb "github.com/brotherlogic/mdb/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	LOWER = 1
	UPPER = 256

	machinesFound = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "mdb_machine_count",
	})

	lookupError = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "mdb_lookup_error",
	}, []string{"error"})
)

func (s *Server) FillDB() error {
	for i := LOWER; i <= UPPER; i++ {
		ipv4 := fmt.Sprintf("192.168.86.%v", i)
		machine, err := s.lookupv4str(ipv4)
		if err != nil {
			lookupError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)}).Inc()
		} else {
			s.machines = append(s.machines, machine)
		}
	}

	machinesFound.Set(float64(len(s.machines)))

	return nil
}

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func strToIpv4(ipv4 string) uint32 {
	return binary.BigEndian.Uint32(net.ParseIP(ipv4))
}

func (s *Server) lookupv4(ipv4 uint32) (*pb.Machine, error) {
	return s.lookupv4str(ipv4ToString(ipv4))
}

func (s *Server) lookupv4str(ipv4 string) (*pb.Machine, error) {
	addr, err := net.LookupAddr(ipv4)
	if err != nil {
		return nil, fmt.Errorf("lookupaddr error: (%v) -> %v", ipv4, err)
	}

	log.Printf("FOUND %v -> %v", ipv4, addr)

	return &pb.Machine{
		Ipv4:     strToIpv4(ipv4),
		Hostname: addr[0],
	}, nil
}
