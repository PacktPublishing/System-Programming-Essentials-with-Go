package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

const socketPath = "/tmp/example.sock"

var (
	clients        []net.Conn
	messageHistory []string
	mutex          sync.Mutex
)

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
	defer conn.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		mutex.Lock()
		clients = append(clients, conn)
		mutex.Unlock()

		// Send the message history to the new client
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			removeClient(conn)
			break
		}

		message := string(buffer[:n])
		messageHistory = append(messageHistory, message)
		broadcastMessage(message)
	}
}

func broadcastMessage(message string) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, client := range clients {
		client.Write([]byte(message + "\n"))
	}
}

func removeClient(conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, client := range clients {
		if client == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}
