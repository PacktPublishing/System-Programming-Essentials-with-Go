package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// Handle the error of WebSocket upgrade failure.
			log.Printf("Error upgrading to WebSocket: %v", err)
			return // Stop processing if we can't upgrade to WebSocket.
		}

		go func() {
			defer conn.Close()

			for {
				// Read data sent by the client.
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// Handle the error from reading client data.
					log.Printf("Error reading from WebSocket: %v", err)
					break // Stop the loop if there's an error reading from the client.
				}

				// Echo the message back to the client.
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					// Handle the error from writing to the WebSocket.
					log.Printf("Error writing to WebSocket: %v", err)
					break // Stop the loop if there's an error writing to the client.
				}
			}

			// Outside the loop, close the connection
			log.Println("Closing WebSocket connection")
			conn.Close()
		}()
	}))
}
