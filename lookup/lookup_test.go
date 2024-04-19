package loookup

import "testing"

func TestIpv4ToStringLookup(t *testing.T) {
	ip := "192.168.86.22"
	num := strToIpv4(ip)
	if num != 3232257558 {
		t.Errorf("Bad ip translation %v -> %v", ip, num)
	}
}
