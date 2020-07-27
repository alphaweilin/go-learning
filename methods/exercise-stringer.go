package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (v IPAddr) String() string {
	var s []string
	for idx, val := range v {
		if idx < len(v)-1 {
			s = append(s, strconv.Itoa(int(val)), ".")
		} else {
			s = append(s, strconv.Itoa(int(val)))
		}
	}
	return strings.Join(s, "")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
