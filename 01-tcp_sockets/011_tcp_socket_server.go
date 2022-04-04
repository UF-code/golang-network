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
	//
	SocketServer(connection_type, connection_addr)
}

func SocketServer(cnx_type string, cnx_addr string) {
	//
	fmt.Printf("Starting %v Server on %v\n", cnx_type, cnx_addr)

	//
	listen, err := net.Listen(cnx_type, cnx_addr)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		os.Exit(1)
	}

	defer listen.Close()

	connection, err := listen.Accept()
	if err != nil {
		fmt.Println("Error Connecting:", err.Error())
		return
	}

	//
	fmt.Printf("Client %v Connected\n", connection.RemoteAddr().String())

	//
	handleConnection(connection)

}

//
func handleConnection(cnx net.Conn) {
	//
	buffer, err := bufio.NewReader(cnx).ReadBytes('\n')
	if err != nil {
		fmt.Printf("Client %v Left\n", cnx.RemoteAddr().String())
		cnx.Close()
		return
	}

	//
	log.Printf("Client %v Message: %v", cnx.RemoteAddr().String(), string(buffer[:len(buffer)-1]))

	//
	cnx.Write(buffer)

	//
	// handleConnection(cnx)

}
