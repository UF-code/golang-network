package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	hostName := "localhost"
	portNum := "8080"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	reader := bufio.NewReader(os.Stdin)

	// The Message We are going to Send to the Server
	fmt.Print("Text to be Sent: ")
	input, _ := reader.ReadString('\n')
	message := []byte(input)

	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}

	// receive message from server
	// receive message from server
	// receive message from server
	// buffer := make([]byte, 1024)
	// n, addr, err := conn.ReadFromUDP(buffer)

	// fmt.Println("UDP Server : ", addr)
	// fmt.Println("Received from UDP server : ", string(buffer[:n]))

}
