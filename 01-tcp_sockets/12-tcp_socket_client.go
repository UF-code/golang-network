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
	connection_type = "tcp"
)

func main() {
	// Start the client and connect to the server.
	fmt.Printf("Connecting to %v Server %v\n",
		connection_type, connection_addr)

	connection, err := net.Dial(connection_type, connection_addr)
	if err != nil {
		fmt.Println("Error Connecting: ", err.Error())
		os.Exit(1)
	}

	// create new reader from stdin
	reader := bufio.NewReader(os.Stdin)

	// run loop forever, until exit
	for {
		// prompting message
		fmt.Print("Text to send: ")

		// read in input until newline, enter key
		input, _ := reader.ReadString('\n')

		// send to socket connection
		connection.Write([]byte(input))

		// listen for relay
		message, _ := bufio.NewReader(connection).ReadString('\n')

		// print server relay
		log.Print("server relay: " + message)
	}
}
