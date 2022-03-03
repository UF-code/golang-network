package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connection_host = "localhost"
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "udp"
)

func main() {
	udpSocketClient(connection_type, connection_addr)
}

func udpSocketClient(cnx_type string, cnx_addr string) {
	// Prompting the Connection of Client
	fmt.Printf("Connecting to %v Server %v \n", cnx_type, cnx_addr)

	udp_addr, err := net.ResolveUDPAddr(cnx_type, cnx_addr)
	if err != nil {
		log.Fatal(err)
	}

	udpDial, err := net.DialUDP(cnx_type, nil, udp_addr)

	log.Printf("Established connection to %s \n", cnx_addr)
	log.Printf("Remote UDP address : %s \n", udpDial.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", udpDial.LocalAddr().String())

	defer udpDial.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to Send: ")

		input, _ := reader.ReadString('\n')

		udpDial.Write([]byte(input))

		message, _ := bufio.NewReader(udpDial).ReadString('\n')

		log.Print("Server relay: ", message)
	}

}
