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
	connection_host = "localhost"
	connection_port = "8080"
	connection_addr = connection_host + ":" + connection_port
	connection_type = "tcp"
)

func main() {
	// Connecting to the TCP Socket Server
	SocketClient(connection_type, connection_addr)

}

func SocketClient(cnx_type string, cnx_addr string) {
	// Prompting the Connection of Client
	fmt.Printf("Connecting to %v Server %v \n", cnx_type, cnx_addr)

	// Connecting to TCP Socket Server
	dial, err := net.Dial(cnx_type, cnx_addr)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		os.Exit(1)
	}

	// End of the Program Dialer Will Close
	defer dial.Close()

	// Creating New Reader From Stdin
	reader := bufio.NewReader(os.Stdin)

	// Sending Message to the Server Forever Until It's Interrupted
	for {
		// Prompting Message
		fmt.Print("Text to Send: ")

		// Read Input Until Newline, 'enter key'
		input, _ := reader.ReadString('\n')

		// Send Socket Connection
		dial.Write([]byte(input))

		// Listen For Relay
		message, _ := bufio.NewReader(dial).ReadString('\n')

		// Prompt the Server Relay
		log.Print("Server relay: ", message)
	}
}
