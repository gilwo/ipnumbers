package ipnumbers

import (
	"fmt"
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
}
