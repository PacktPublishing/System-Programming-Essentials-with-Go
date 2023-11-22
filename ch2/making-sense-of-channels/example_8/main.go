package main

import "fmt"

func main() {
	balls := make(chan string)

	go throwBalls("red", balls)
	go throwBalls("green", balls)

	for color := range balls {
		fmt.Println(color, "received!")
	}
}

func throwBalls(color string, balls chan string) {
	fmt.Printf("throwing the %s ball\n", color)
	balls <- color
}
