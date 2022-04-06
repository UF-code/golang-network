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
	fmt.Printf("Starting %v Server on %v\n",
		connection_type, connection_addr)

	listen, err := net.Listen(connection_type, connection_addr)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		os.Exit(1)
	}
	// Close the listener when the applicationn close
	defer listen.Close()

	// run loop forever, until exit
	for {
		// listen for an incoming connection
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error Connecting:", err.Error())
			return
		}
		fmt.Println("Client Connected...")

		fmt.Printf("Client %v Connected\n", connection.RemoteAddr().String())

		// handle connections concurrently in a new goroutine
		handleConnection(connection)
	}

}

// handleConnection handles logic for a single connection request
func handleConnection(conn net.Conn) {
	// buffer client input until a newline
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// close left clients
	if err != nil {
		fmt.Println("client left")
		conn.Close()
		return
	}

	// Print response message, stripping newline character
	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	// send response message to the client
	conn.Write(buffer)

	// restart the process
	handleConnection(conn)

}
