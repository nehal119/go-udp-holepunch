package main

import (
	"fmt"
	"net"
)

// type clientType map[string]bool

// var clients = clientType{}

func Server1() {
	address := "0.0.0.0:9595"

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on", addr)

	for {
		fmt.Println("Got incoming something")
		// Listen for messages
		msg := make([]byte, 1024)
		n, remote, err := conn.ReadFromUDP(msg)
		if err != nil {
			fmt.Println("conn.ReadFromUDP error", err.Error())
			continue
		}
		message := string(msg[0:n])
		fmt.Println("Got incoming message", message)

		if message == "register" {
			// Store this ip
			clients[remote.String()] = true

			for client := range clients {
				ipAddress := clients.keys(client)
				if len(ipAddress) > 0 {
					// Write message
					addr, _ := net.ResolveUDPAddr("udp", client)
					conn.WriteToUDP([]byte(ipAddress), addr)
					fmt.Println("Message written to", client)
				}
			}
		}
	}
}
