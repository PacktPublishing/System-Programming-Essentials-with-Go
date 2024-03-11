package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, TLS!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
