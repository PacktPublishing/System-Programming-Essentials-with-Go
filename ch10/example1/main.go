package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://pokeapi.co/api/v2/pokemon/ditto"

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
