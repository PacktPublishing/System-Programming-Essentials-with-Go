package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server at the UNIX socket
	conn, err := net.Dial("unix", "/tmp/example.sock")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	// Send a message
	_, err = conn.Write([]byte("Hello UNIX socket!\n"))

	if err != nil {
		fmt.Println("Error writing to socket:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading from socket:", err)
		return
	}

	fmt.Println("Server response:", string(buffer[:n]))
}
