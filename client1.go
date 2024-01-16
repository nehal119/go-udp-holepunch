package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func Client1(clientAddr string) {
	// clientAddr := ":4545"
	remoteAddr := "127.0.0.1:9595"
	client, _ := net.ResolveUDPAddr("udp", clientAddr)
	remote, _ := net.ResolveUDPAddr("udp", remoteAddr)

	conn, err := net.ListenUDP("udp", client)
	if err != nil {
		panic(err)
	}

	// Send register message
	go func() {
		conn.WriteTo([]byte("register"), remote)
		fmt.Println("Client registered", client)
	}()

	// Transmit messages
	for {
		// Read and print message received
		msg := make([]byte, 1024)
		// n, err := conn.Read(msg)
		n, remote, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}
		messages := string(msg[0:n])
		fmt.Println("Connection recieved message", messages, "from", remote.String())
		// Make sure it is the first time, otherwise we will get into a loop
		if messages == "Hello" {
			continue
		}
		// Check if the current user is sender receiver then send a response message
		for _, cl := range strings.Split(messages, ",") {
			if cl != client.String() {
				// Send message
				go chat(conn, cl)
			}
		}
	}
}

func chat(conn *net.UDPConn, cl string) {
	addr, _ := net.ResolveUDPAddr("udp", cl)
	for {
		// Send message
		conn.WriteTo([]byte("Hello"), addr)
		fmt.Println("Connection sent message to", cl)
		time.Sleep(5 * time.Second)
	}
}
