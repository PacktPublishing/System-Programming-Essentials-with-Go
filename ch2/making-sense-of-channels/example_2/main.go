package main

import "fmt"

func main() {
	c := make(chan string)
	fmt.Println(<-c)
}
