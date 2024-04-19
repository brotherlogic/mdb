package loookup

import (
	"context"
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

func getMacAddress(addr string) (string, string) {
	out, err := exec.Command("/usr/bin/nmap", addr).CombinedOutput()
	if err != nil {
		return fmt.Sprintf("unable to nmap: %v -> %v", err, string(out)), ""
	}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "MAC") && !strings.Contains(line, "Unable") {
			return strings.Split(line, " ")[2], strings.Join(strings.Split(line, " ")[3:], " ")
		}
	}

	return fmt.Sprintf("Unable to find mac address: %v", string(out)), ""
}

func mergeInto(machines []*pb.Machine, machine *pb.Machine) []*pb.Machine {
	for _, exMachine := range machines {
		if exMachine.Controller == machine.Controller && exMachine.Hostname == machine.Hostname {
			exMachine.Ipv4 = machine.Ipv4
			exMachine.Mac = machine.Mac
			return machines
		}
	}

	return append(machines, machine)
}

func FillDB(ctx context.Context, mdb *pb.Mdb) error {
	for i := LOWER; i <= UPPER; i++ {
		ipv4 := fmt.Sprintf("192.168.86.%v", i)
		machine, err := lookupv4str(ipv4)
		if err != nil {
			lookupError.With(prometheus.Labels{"error": fmt.Sprintf("%v", err)}).Inc()
		} else {
			mac, thing := getMacAddress(ipv4)
			log.Printf("%v -> %v, %v", ipv4, mac, thing)
			machine.Mac = mac
			machine.Controller = thing
			mdb.Machines = mergeInto(mdb.Machines, machine)
		}
	}

	log.Printf("Setting machine count: %v", len(mdb.Machines))
	machinesFound.Set(float64(len(mdb.Machines)))

	return nil
}

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func strToIpv4(ipv4 string) uint32 {
	return binary.BigEndian.Uint32(net.ParseIP(ipv4)[12:16])
}

func lookupv4(ipv4 uint32) (*pb.Machine, error) {
	return lookupv4str(ipv4ToString(ipv4))
}

func lookupv4str(ipv4 string) (*pb.Machine, error) {
	addr, err := net.LookupAddr(ipv4)
	if err != nil {
		return nil, fmt.Errorf("lookupaddr error: (%v) -> %v --> %v", ipv4, err, strToIpv4(ipv4))
	}

	log.Printf("FOUND %v -> %v --> %v", ipv4, addr, strToIpv4(ipv4))

	return &pb.Machine{
		Ipv4:     strToIpv4(ipv4),
		Hostname: addr[0],
	}, nil
}
