package ipnumbers

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {

	strIPA := []string{"192.168.1.1", "1.1.1.2/24", "dead::beaf", "2001::1/64"}
	for _, strIP := range strIPA {
		high, low, _ := IPtouint64(strIP)
		fmt.Printf("ip: %s, representation in numbers: [%v, %v]\n", strIP, high, low)
		ipstr, ip := Uint64toip(high, low)
		fmt.Printf("convert back to ip from high: %v, low: %v - ip : %v\n", high, low, ipstr)
		if ip.String() != strings.Split(strIP, "/")[0] {
			t.Errorf("original ip (%s) and after convertion (%s) do not match\n", strIP, ipstr)
		}
	}

	strIPs := []string{"1.2.3.4", "2001::1"}
	for _, e := range strIPs {
		ip := net.ParseIP(e)
		high1, low1 := NetIPtouint64(&ip)
		high2, low2, err := IPtouint64(e)
		if err != nil {
			t.Errorf("failed converting address %s to number", e)
		}
		if high1 != high2 || low1 != low2 {
			t.Errorf("mismatch convertion of %s to number (string and net.IP)", e)
		}
	}
}
