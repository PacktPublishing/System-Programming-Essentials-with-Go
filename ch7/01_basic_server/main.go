package main

import (
	"fmt"
	"net"
	"os"
)

const socketPath = "/tmp/example.sock"

func main() {
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	conn.Close()
}
