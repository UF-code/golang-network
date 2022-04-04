package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Setting up The Connection Variables
const (
	connection_host = "localhost" // 127.0.0.1
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "tcp"
)

func main() {
	// Creating TCP Socket Server
	SocketServer(connection_type, connection_addr)
}

func SocketServer(cnx_type string, cnx_addr string) {
	// Prompting the Socket Type and Socket Address
	fmt.Printf("Starting %v Server on %v\n", cnx_type, cnx_addr)

	// Creating a Listener Who Listens and Answers the Connected Client
	listen, err := net.Listen(cnx_type, cnx_addr)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		os.Exit(1)
	}

	// End of the Program Listener Will Close
	defer listen.Close()

	// Accepts the Client Who wants to Connect
	connection, err := listen.Accept()
	if err != nil {
		fmt.Println("Error Connecting:", err.Error())
		return
	}

	// Prompting Client Address
	fmt.Printf("Client %v Connected\n", connection.RemoteAddr().String())

	// Handling the Connection
	handleConnection(connection)
}

//
func handleConnection(cnx net.Conn) {
	// Creating a Buffer to Read Received Bytes From Connection
	message, err := bufio.NewReader(cnx).ReadBytes('\n')
	if err != nil {
		fmt.Printf("Client %v Left\n", cnx.RemoteAddr().String())
		cnx.Close()
		return
	}

	// Prompting Converted Message From Client Who Send the Message
	log.Printf("Client %v Message: %v", cnx.RemoteAddr().String(), string(message[:len(message)-1]))

	// Writing Message
	cnx.Write(message)
}
