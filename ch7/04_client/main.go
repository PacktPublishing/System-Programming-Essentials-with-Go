package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

const socketPath = "/tmp/chat.sock"

func main() {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	// Listen for messages from the server
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println("Message from server:", scanner.Text())
		}
	}()

	// Send messages to the server
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter message:")
	for scanner.Scan() {
		message := scanner.Text()
		conn.Write([]byte(message))
	}

	wg.Wait()
}
