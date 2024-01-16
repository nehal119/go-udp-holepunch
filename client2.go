package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

// var started = false

func Client2(clientAddr, clientName string) {
	// clientAddr := ":4545"
	// remoteAddr := "127.0.0.1:9595"
	client, _ := net.ResolveUDPAddr("udp", clientAddr)
	// remote, _ := net.ResolveUDPAddr("udp", remoteAddr)

	conn, err := net.ListenUDP("udp", client)
	if err != nil {
		panic(err)
	}

	// Send register message
	// Get other person IP and Port and send ours
	resp, err := sendAndgetIP(clientName, clientAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name", resp.Name)
	fmt.Println("Address", resp.Addr)

	// Start sending mesages
	// if !started {
	go chat2(conn, resp.Addr)
	// 	started = true
	// }

	i := 0

	for {
		if i > 10 {
			break
		}
		// Read and print message received
		msg := make([]byte, 1024)
		// n, err := conn.Read(msg)
		n, remoteIP, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}
		messages := string(msg[0:n])
		fmt.Println("Connection recieved message", messages, "from", remoteIP.String())
		i++
		// Make sure it is the first time, otherwise we will get into a loop
		// if messages == "Hello" {
		// 	continue
		// }
	}

	// Send message
	// addr, _ := net.ResolveUDPAddr("udp", resp.Addr)
	// conn.WriteTo([]byte("Hello"), addr)
	// fmt.Println("Connection sent message to", resp.Addr)
	// time.Sleep(5 * time.Second)
}

func chat2(conn *net.UDPConn, cl string) {
	addr, _ := net.ResolveUDPAddr("udp", cl)
	for {
		// Send message
		conn.WriteTo([]byte("Hello"), addr)
		fmt.Println("Connection sent message to", cl)
		time.Sleep(5 * time.Second)
	}
}

func sendAndgetIP(clientName, clientAddr string) (Resp, error) {
	url := "http://127.0.0.1:33333/register"
	method := "POST"

	payload := strings.NewReader(`{"client": "` + clientName + `", "address": "` + clientAddr + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return Resp{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Resp{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Resp{}, err
	}
	fmt.Println(string(body))
	// Response body has name, address and port. Unmarshal it

	var response Resp
	json.Unmarshal(body, &response)
	fmt.Println(response.Name, response.Addr)

	return response, nil
}

type Resp struct {
	Name string `json:"name"`
	Addr string `json:"address"`
}
