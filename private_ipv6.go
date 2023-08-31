package main

// create a random IPv6 private CIDR (/64)
//
// Usage: ./private_ipv6.py
//
// Author: Guy Davies
// Email: aguydavies@gmail.com

import (
	"fmt"
	"strings"

	"github.com/guydavies/hexmanip"
)

func main() {
	// construct random private IPv6 CIDR block
	//
	// First construct the first group which must begin with 'fd' and append a
	// random byte
	//
	// Then iterate three times to construct three more random hex blocks of 4
	// hex digits (two bytes) each
	var max_hex_value_as_int int = 65535
	var hex_first_byte string = "fd"
	var hex_two_bytes []string
	var addr_all_groups []string
	hex_two_bytes = append(hex_two_bytes, hex_first_byte)
	var hex_byte string = hexmanip.Generate_hex(255)
	hex_two_bytes = append(hex_two_bytes, hex_byte)
	var hex_two_byte_str string = strings.Join(hex_two_bytes, "")
	addr_all_groups = append(addr_all_groups, hex_two_byte_str)

	for i := 0; i < 3; i++ {
		addr_all_groups = append(addr_all_groups, hexmanip.Generate_hex(max_hex_value_as_int))
	}

	addr_all_groups = append(addr_all_groups, ":/64")
	var ipv6_address string = strings.Join(addr_all_groups[:], ":")
	fmt.Printf(ipv6_address)
	fmt.Println("")
}
