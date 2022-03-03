package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection_host := "localhost"
	connection_port := "8080"

	connection_addr := connection_host + ":" + connection_port

	RemoteAddr, err := net.ResolveUDPAddr("udp", connection_addr)

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", connection_addr)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	message := []byte("Go is Fun!")

	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}

	// receive message from server
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP Server : ", addr)
	fmt.Println("Received from UDP server : ", string(buffer[:n]))

}
