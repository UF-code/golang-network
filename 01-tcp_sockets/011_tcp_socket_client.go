package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connection_host = "localhost" // 127.0.0.1
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "tcp"
)

func main() {
	//
	fmt.Printf("Connecting to %v Server %v \n", connection_type, connection_addr)

	connection, err := net.Dial(connection_type, connection_addr)
	if err != nil {
		fmt.Println("Error Connecting: ", err.Error())
		os.Exit(1)
	}
	defer connection.Close()

	//
	reader := bufio.NewReader(os.Stdin)

	//
	fmt.Print("Text to send: ")

	//
	input, _ := reader.ReadString('\n')

	//
	connection.Write([]byte(input))

	//
	// message, _ := bufio.NewReader(connection).ReadString('\n')

	//
	// log.Print("Server Relay", message)
}
