package main

import (
	"os"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "c1":
		Client1(":4545")
	case "c2":
		Client1(":4546")
	case "s":
		Server1()
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
