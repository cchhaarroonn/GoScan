package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var address string
	ports := [6]string{"21", "22", "23", "80", "443", "25565"}
	var opened int
	var timeout int

	fmt.Print("[*] Enter address you want to scan: ")
	fmt.Scanln(&address)
	fmt.Print("[*] Enter timeout for each port: ")
	fmt.Scanln(&timeout)
	fmt.Print("\n")

	results := make(chan string)

	for _, port := range ports {
		go func(port string) {
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(address, port), time.Duration(timeout)*time.Second)
			if err != nil {
				results <- fmt.Sprintf("[X] Port %s is closed on address %s\n", port, address)
			} else {
				results <- fmt.Sprintf("[*] Port %s is open on address %s\n", port, address)
				conn.Close()
			}
		}(port)
	}

	for range ports {
		fmt.Print(<-results)
		opened++
	}
	fmt.Printf("\n[*] Scan completed. Found %d port(s) open.\n", opened)
}
