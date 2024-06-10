package main

import (
	"fmt"
	"net"
)

func main() {
	// Start listening for connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Accept connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Printf("write error: %v\n", err)
		}
	}
}
