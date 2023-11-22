package main

import "fmt"

func main() {
	balls := make(chan string)
	go throwBalls("red", balls)
	fmt.Println(<-balls, "received!")
}

func throwBalls(color string, balls chan string) {
	fmt.Printf("throwing the %s ball\n", color)
	balls <- color
}
