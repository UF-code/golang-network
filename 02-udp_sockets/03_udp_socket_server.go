package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
)

// Setting up The Connection Variables
const (
	connection_host = "localhost"
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "udp"
)

func main() {
	// Number of Goroutines Currently Occupied
	fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())

	// Creating the TCP Socket Server
	udpSocketServer(connection_type, connection_addr)
}

func udpSocketServer(cnx_type string, cnx_addr string) {
	// Prompting the Socket Type and Socket Address
	fmt.Printf("Starting %v Server on %v\n",
		cnx_type, cnx_addr)

	//
	udp_addr, err := net.ResolveUDPAddr("udp4", cnx_addr)
	if err != nil {
		log.Fatal(err)
	}

	//
	listen, err := net.ListenUDP(cnx_type, udp_addr)
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		_, connection, err := listen.ReadFromUDP()
	}

}

func handleUDPConnection() {
	fmt.Println("udp socket server")
}
