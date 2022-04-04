package main

import (
	"bufio"
	"fmt"
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
	// Connecting to The TCP Socket Server
	SocketClient(connection_type, connection_addr)
}

func SocketClient(cnx_type string, cnx_addr string) {
	// Prompting the Connection of Client
	fmt.Printf("Connecting to %v Server %v \n", cnx_type, cnx_addr)

	// Connecting to The TCP Socket Server
	connection, err := net.Dial(cnx_type, cnx_addr)
	if err != nil {
		fmt.Println("Error Connecting: ", err.Error())
		os.Exit(1)
	}
	// End of the Program Connection Will Close
	defer connection.Close()

	// Creating New Reader From Stdin
	reader := bufio.NewReader(os.Stdin)

	// The Message We are going to Send to the Server
	fmt.Print("Text to be Sent: ")
	input, _ := reader.ReadString('\n')

	// Sending Message to Server
	connection.Write([]byte(input))
}
