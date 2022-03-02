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
	connection_type = "tcp"
)

func main() {
	//

}

func SocketClient(cnx_type string, cnx_addr string) {
	//
	fmt.Printf("Connecting to %v Server %v \n", cnx_type, cnx_addr)

	connection, err := net.Dial(cnx_type, cnx_addr)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		os.Exit(1)
	}
	//
	defer listen.Close()

	//
	reader := bufio.NewReader(os.Stdin)

	//
	for {
		//
		fmt.Print("Text to Send: ")

		//
		input, _ := reader.ReadString('\n')

		//
		connection.Write([]byte(input))

		//
		message, _ := bufio.NewReader(connection).ReadString('\n')

		//
		log.Print("Server relay: ", message)
	}
}
