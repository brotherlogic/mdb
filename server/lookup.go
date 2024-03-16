package server

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"

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

func (s *Server) getMacAddress(addr string) (string, string) {
	out, err := exec.Command("/usr/bin/nmap", addr).CombinedOutput()
	if err != nil {
		return fmt.Sprintf("unable to nmap: %v -> %v", err, string(out)), ""
	}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "MAC") {
			return strings.Split(line, " ")[2], strings.Join(strings.Split(line, " ")[3:], " ")
		}
	}

	return fmt.Sprintf("Unable to find mac address: %v", string(out)), ""
}

func (s *Server) FillDB() error {
	for i := LOWER; i <= UPPER; i++ {
		ipv4 := fmt.Sprintf("192.168.86.%v", i)
		machine, err := s.lookupv4str(ipv4)
		if err != nil {
			lookupError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)}).Inc()
		} else {
			mac, thing := s.getMacAddress(ipv4)
			log.Printf("%v -> %v, %v", ipv4, mac, thing)
			machine.Mac = mac
			machine.Controller = thing
			s.machines = append(s.machines, machine)
		}
	}

	log.Printf("Setting machine count: %v", len(s.machines))
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
