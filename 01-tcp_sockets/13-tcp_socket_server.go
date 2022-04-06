package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

// Setting up The Connection Variables
const (
	connection_host = "localhost" // 127.0.0.1
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "tcp"
)

func main() {
	// Number of Goroutines Currently Occupied
	fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())

	// Creating the TCP Socket Server
	SocketServer(connection_type, connection_addr)
}

func SocketServer(cnx_type string, cnx_addr string) {
	// Prompting the Socket Type and Socket Address
	fmt.Printf("Starting %v Server on %v\n",
		cnx_type, cnx_addr)

	// Creating a Listener Who Listens the Connected Client
	listen, err := net.Listen(cnx_type, cnx_addr)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		os.Exit(1)
	}

	// End of the Program Listener Will Close
	defer listen.Close()

	// Getting Data From Clients Forever Until It's Interrupted
	for {
		// Accepts the Connection
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error Connecting:", err.Error())
			return
		}
		// Prompting Clients Address
		fmt.Printf("Client %v Connected\n", connection.RemoteAddr().String())

		// Running Concurrent Connection Handler
		go handleConnection(connection)

		// Number of Goroutines Currently Occupied
		fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())

	}
}

func handleConnection(cnx net.Conn) {
	// Creating a Buffer to Read Received Btyes From Connection
	buffer, err := bufio.NewReader(cnx).ReadBytes('\n')

	// When the Client Left It Causes Error and We Prompting That Message
	if err != nil {
		fmt.Printf("Client %v Left\n", cnx.RemoteAddr().String())
		cnx.Close()
		return
	}

	// Prompting Converted Message From Client Who Send the Message
	log.Printf("Client %v Message: %v", cnx.RemoteAddr().String(),
		string(buffer[:len(buffer)-1]))

	// Write to the Socket
	cnx.Write(buffer)

	// Restart the Handle Connection
	handleConnection(cnx)
}
