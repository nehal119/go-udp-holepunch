package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("RENDEZVOUS_SERVER", os.Getenv("RENDEZVOUS_SERVER"))
	// ./hp CLIENT_NAME LISTENER_ADDRESS PUBLIC_IP PORT
	if len(os.Args) < 5 {
		fmt.Println("Usage: ./hp client_name listener_address my_public_ip port")
		fmt.Println("Example: ./hp A 0.0.0.0 56.33.99.88 4545")
		return
	}
	cmd := os.Args[1]
	clientListenerAddr := os.Args[2]
	clientPublicIp := os.Args[3]
	port := os.Args[4]
	switch cmd {
	// case "c1":
	// 	Client1(":4545")
	// case "c2":
	// 	Client1(":4546")
	case "A":
		Client2("A", clientListenerAddr+":"+port, clientPublicIp+":"+port)
	case "B":
		Client2("B", clientListenerAddr+":"+port, clientPublicIp+":"+port)
		// case "s":
		// 	Server1()
	}
}

// func main() {
// 	cmd := os.Args[1]
// 	switch cmd {
// 	case "c":
// 		Client()
// 	case "s":
// 		Server()
// 	}
// }
