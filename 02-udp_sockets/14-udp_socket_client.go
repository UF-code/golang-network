package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connection_host = "localhost" // 127.0.0.1
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "udp"
)

func main() {
	udpSocketClient(connection_type, connection_addr)
}

func udpSocketClient(cnx_type string, cnx_addr string) {
	//
	fmt.Printf("Connecting to %v Server %v \n", cnx_type, cnx_addr)

	//
	RemoteAddr, err := net.ResolveUDPAddr(cnx_type, cnx_addr)

	//
	cnx, err := net.DialUDP(cnx_type, nil, RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer cnx.Close()

	//
	reader := bufio.NewReader(os.Stdin)

	//
	fmt.Print("Text to be Sent: ")
	input, _ := reader.ReadString('\n')
	message := []byte(input)

	_, err = cnx.Write(message)
	if err != nil {
		log.Println(err)
	}

}
