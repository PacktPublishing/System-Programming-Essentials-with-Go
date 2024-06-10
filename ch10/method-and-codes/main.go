package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/resource", resourceHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HTTP verbs and status codes example!")
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Handle GET request
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "Resource fetched successfully")
	case "POST":
		// Handle POST request
		w.WriteHeader(http.StatusCreated) // 201
		fmt.Fprintf(w, "Resource created successfully")
	case "PUT":
		// Handle PUT request
		w.WriteHeader(http.StatusAccepted) // 202
		fmt.Fprintf(w, "Resource updated successfully")
	case "DELETE":
		// Handle DELETE request
		w.WriteHeader(http.StatusNoContent) // 204
	default:
		// Handle unknown methods
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintf(w, "Method not allowed")
	}
}
