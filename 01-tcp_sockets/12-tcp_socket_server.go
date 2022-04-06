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

	// Getting Data From Clients Forever Until It's Interrupted
	for {
		// Listen for an incoming connection
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error Connecting:", err.Error())
			return
		}
		fmt.Println("Client Connected...")

		fmt.Printf("Client %v Connected\n", connection.RemoteAddr().String())

		// Handling Connection
		handleConnection(connection)
	}

}

// HandleConnection handles logic for a single connection request
func handleConnection(cnx net.Conn) {
	// Buffer client input until a newline
	buffer, err := bufio.NewReader(cnx).ReadBytes('\n')

	// Close left clients
	if err != nil {
		fmt.Printf("Client %v Left\n", cnx.RemoteAddr().String())
		cnx.Close()
		return
	}

	// Print response message, stripping newline character
	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	// Send response message to the client
	cnx.Write(buffer)

	// Restart the process
	handleConnection(cnx)
}
