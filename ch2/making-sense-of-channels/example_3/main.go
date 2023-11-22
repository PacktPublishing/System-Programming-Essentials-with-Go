package main

import "fmt"

func main() {
	c := make(chan string)
	c <- "message"   // Sending
	fmt.Println(<-c) // Receiving
}
