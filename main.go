package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	var address string
	ports := [6]int{21, 22, 23, 80, 443, 25565}
	var opened int
	var timeout int

	fmt.Print("[*] Enter address you want to scan: ")
	fmt.Scanln(&address)
	fmt.Print("[*] Enter timeout for each port: ")
	fmt.Scanln(&timeout)
	fmt.Print("\n")

	for _, port := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(address, strconv.Itoa(port)), time.Duration(timeout)*time.Second)
		if err != nil {
			fmt.Printf("[X] Port %d is closed on address %s\n", port, address)
		} else {
			fmt.Printf("[*] Port %d is open on address %s\n", port, address)
			opened++
			conn.Close()
		}
	}
	fmt.Printf("[*] Scan completed. %d port(s) open.\n", opened)
}
