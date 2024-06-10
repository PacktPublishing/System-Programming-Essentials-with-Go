package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	socketPath := "/tmp/go-server.sock"
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal("Listen (UNIX socket):", err)
	}
	defer listener.Close()
	log.Println("Server is listening on", socketPath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error writing response:", err)
		}
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		log.Println("Shutting down gracefully...")
		listener.Close()
		os.Remove(socketPath)
		os.Exit(0)
	}()

	err = http.Serve(listener, nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("HTTP server error:", err)
	}
}
