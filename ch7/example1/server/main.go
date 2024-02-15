package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Define the path for the UNIX socket
	socketPath := "/tmp/example.sock"
	// Clean up before start
	if err := os.Remove(socketPath); err != nil && !os.IsNotExist(err) {
		log.Printf("Error removing socket file: %v", err)
		return
	}

	// Listen on the UNIX socket
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Printf("Error listening: %v", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening on", socketPath)

	// Set up a signal handler to gracefully shut down the server
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		fmt.Println("Received termination signal. Shutting down gracefully...")
		listener.Close()
		os.Remove(socketPath)
		os.Exit(0)
	}()

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))

	// Simulate a response back to the client
	response := []byte("Message received successfully\n")

	_, err = conn.Write(response)

	if err != nil {
		log.Printf("Error writing response to connection: %v", err)
		return
	}
}
