package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	ctx := context.Background()
	// Connect to the WebSocket server
	conn, _, _, err := ws.DefaultDialer.Dial(ctx, "ws://localhost:8080")
	if err != nil {
		fmt.Printf("Error connecting to WebSocket server: %v\n", err)
		return
	}
	defer conn.Close()

	// Send a message to the server
	message := []byte("Hello, server!")
	err = wsutil.WriteClientMessage(conn, ws.OpText, message)
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
		return
	}

	// Read the server's response
	response, _, err := wsutil.ReadServerData(conn)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Received from server: %s\n", response)

	// Keep the client running until the user decides to exit
	fmt.Println("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
