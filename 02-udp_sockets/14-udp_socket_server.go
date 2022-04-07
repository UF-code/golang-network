package main

import (
	"fmt"
	"log"
	"net"
)

const (
	connection_host = "localhost" // 127.0.0.1
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "udp"
)

func main() {
	//
	udpSocketServer(connection_type, connection_addr)
}

func udpSocketServer(cnx_type string, cnx_addr string) {

	//
	fmt.Printf("Starting %v Server on %v\n", cnx_type, cnx_addr)

	//
	udpAddr, err := net.ResolveUDPAddr(cnx_type, cnx_addr)
	if err != nil {
		log.Fatal(err)
	}

	//
	cnx, err := net.ListenUDP(cnx_type, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer cnx.Close()

	handleConnection(cnx)
}

func handleConnection(cnx *net.UDPConn) {
	//
	buffer := make([]byte, 1024)

	//
	msg, addr, err := cnx.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
		cnx.Close()
	}

	fmt.Printf("UDP Client %v's Message:  %v\n", addr, string(buffer[:msg]))
}
