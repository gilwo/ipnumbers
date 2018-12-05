package main

import (
	"fmt"
	"github.com/gilwo/ipnumbers"
	"github.com/jessevdk/go-flags"
	"os"
)

var Opts struct {
	NumIP []uint64 `short:"n" long:"number" description:"64 bit number representing ip number (1 for ipv4, 2 for ipv6)"`
	StrIP string   `short:"s" long:"string" description:"string representing ip number"`
}

func main() {
	if _, err := flags.NewParser(&Opts, flags.Default).Parse(); err != nil {
		if e, ok := err.(*flags.Error); ok {
			switch e.Type {
			case flags.ErrHelp:
				//fmt.Println("help requested, existing")
				os.Exit(0)
			case flags.ErrInvalidChoice:
				fmt.Println("invalid choice")
			default:
				fmt.Printf("error parsing opts: %v\n", e.Type)
			}
		} else {
			fmt.Printf("unknown error: %v\n", err)
		}
		os.Exit(1)
	} else if len(Opts.NumIP) == 0 && len(Opts.StrIP) == 0 {
		fmt.Printf("not enough arguments\n")
		flags.NewParser(&Opts, flags.Default).WriteHelp(os.Stderr)
		os.Exit(1)
	}

	fmt.Printf("args: %q\n", os.Args)

	if len(Opts.NumIP) > 1 {
		high, low := Opts.NumIP[0], Opts.NumIP[1]
		fmt.Printf("ipnum - high: %d, low: %d\n", high, low)
		ipstr, _ := ipnumbers.Uint64toip(high, low)
		fmt.Printf("ip: %s\n", ipstr)
	} else if len(Opts.NumIP) > 0 {
		low := Opts.NumIP[0]
		fmt.Printf("ipnum - low: %d\n", low)
		ipstr, _ := ipnumbers.Uint64toip(0, low)
		fmt.Printf("ip: %s\n", ipstr)
	} else if len(Opts.StrIP) > 0 {
		ip := Opts.StrIP
		fmt.Printf("ipnum - str: %s\n", ip)
		high, low, err := ipnumbers.IPtouint64(ip)
		if err != nil {
			fmt.Errorf("conversion error: %v\n", err)
		} else {
			fmt.Printf("ipnum: high: %v, low: %v\n", high, low)
		}
	}
}
