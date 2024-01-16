package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// type clientType map[string]bool
// var clients = clientType{}

func Server2() {
	address := "0.0.0.0:9595"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()

	fmt.Println("Got incoming something")
	// Listen for messages
	reader := bufio.NewReader(conn)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("reader.ReadString error", err.Error())
		return
	}
	message := strings.TrimSpace(msg)
	fmt.Println("Got incoming message", message)

	if message == "register" {
		// Store this ip
		clients[conn.RemoteAddr().String()] = true

		for client := range clients {
			has := clients.keys(client)
			if len(has) > 0 {
				// Write message
				conn.Write([]byte(has + "\n"))
				fmt.Println("Message written to", client)
			}
		}
	}
}
