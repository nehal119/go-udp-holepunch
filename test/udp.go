package main

import (
	"fmt"
	"net"
)

func main() {
	address := "0.0.0.0:9595"

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP server started on", addr)

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", addr.String(), message)
	}
}
