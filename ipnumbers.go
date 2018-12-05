//  convert ipv4 to number representation:
//  ipv4 is 4 octec which can be represented in 32 bit number
//
//  convert ipv6 to number represenation:
//  ipv6 is 16 octec which can be represneted in 2 64 bit number (a bit simpler then one 128 bit number)
//
//  ipv4 example:
//
//  ip: 192.168.1.1
//  number: 192 * (1 << (8*3)) + 168 * (1 << (8*2)) + 1 * (1 << (8*1) + 1 * (1 << (8*0)
//   = 192 * 16777216 + 168 * 65536 + 1 * 256 + 1 * 1
//   = 3232235777
//
//  convert number to ipv4:
//  number = 3232235777
//  ip:
//   first octec = (3232235777 & (255 << (8*3))) / (1 << (8*3)) = 192
//   first octec = (3232235777 & (255 << (8*2))) / (1 << (8*2)) = 168
//   first octec = (3232235777 & (255 << (8*1))) / (1 << (8*1)) = 1
//   first octec = (3232235777 & (255 << (8*0))) / (1 << (8*0)) = 1
package ipnumbers

import (
	"fmt"
	"net"
)

const (
	octecBase8 uint64 = 1 << (8 * iota)
	octecBase7 uint64 = 1 << (8 * iota)
	octecBase6 uint64 = 1 << (8 * iota)
	octecBase5 uint64 = 1 << (8 * iota)
	octecBase4 uint64 = 1 << (8 * iota)
	octecBase3 uint64 = 1 << (8 * iota)
	octecBase2 uint64 = 1 << (8 * iota)
	octecBase1 uint64 = 1 << (8 * iota)
)

const (
	octecMask8 uint64 = 255 << (8 * iota)
	octecMask7 uint64 = 255 << (8 * iota)
	octecMask6 uint64 = 255 << (8 * iota)
	octecMask5 uint64 = 255 << (8 * iota)
	octecMask4 uint64 = 255 << (8 * iota)
	octecMask3 uint64 = 255 << (8 * iota)
	octecMask2 uint64 = 255 << (8 * iota)
	octecMask1 uint64 = 255 << (8 * iota)
)

var octecBaseSlice = []uint64{octecBase1, octecBase2, octecBase3, octecBase4, octecBase5, octecBase6, octecBase7, octecBase8}
var octecMaskSlice = []uint64{octecMask1, octecMask2, octecMask3, octecMask4, octecMask5, octecMask6, octecMask7, octecMask8}

// convert an ip (either ipv4 or ipv6) from string represnetation
// to number representation in 2 uint64 numbers, high and low
//
// invalid ip address will cause an error to return (high and low will be 0)
func IPtouint64(ip string) (high, low uint64, err error) {

	ipaddr, _, err := net.ParseCIDR(ip)
	if err != nil {
		err = nil
		ipaddr = net.ParseIP(ip)
		if ipaddr == nil {
			return high, low, fmt.Errorf("unable to parse ip (%v)", err)
		}
	}

	if ipaddr.To4() != nil {
		ipaddr[10] = 0
		ipaddr[11] = 0
	}

	for i, _ := range octecBaseSlice {
		high += octecBaseSlice[i] * (uint64)(ipaddr[i])
		low += octecBaseSlice[i] * (uint64)(ipaddr[i+8])
	}

	return
}

// convert 2 numbers of uint64 (high and low) to a string representation of the the ip address
// and to a net.IP type object
func Uint64toip(high, low uint64) (ip string, netip net.IP) {
	var bytelow, bytehigh, byteip []byte
	for i, _ := range octecMaskSlice {
		bytelow = append(bytelow, (byte)((octecMaskSlice[i]&low)/octecBaseSlice[i]))
		bytehigh = append(bytehigh, (byte)((octecMaskSlice[i]&high)/octecBaseSlice[i]))
	}

	byteip = append(byteip, bytehigh...)
	byteip = append(byteip, bytelow...)

	if high == 0 {
		byteip[10] = 0xff
		byteip[11] = 0xff
	}

	netip = net.IP(byteip)
	ip = fmt.Sprintf("%s", netip)

	return
}
