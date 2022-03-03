package main

import (
	"fmt"
	"log"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	// write message back to client
	message := []byte("Hello UDP client!")
	_, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}

}

func main() {
	connection_host := "localhost"
	connection_port := "8080"
	connection_addr := connection_host + ":" + connection_port

	udpAddr, err := net.ResolveUDPAddr("udp4", connection_addr)

	if err != nil {
		log.Fatal(err)
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server up and listening on port 8080")

	defer ln.Close()

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln)
	}

}
