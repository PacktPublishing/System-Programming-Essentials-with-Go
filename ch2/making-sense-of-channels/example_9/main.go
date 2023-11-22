package main

import (
	"fmt"
	"sync"
)

func main() {
	balls := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		throwBalls("red", balls)
	}()

	go func() {
		defer wg.Done()
		throwBalls("green", balls)
	}()

	go func() {
		wg.Wait()
		close(balls)
	}()

	for color := range balls {
		fmt.Println(color, "received!")
	}
}

func throwBalls(color string, balls chan string) {
	fmt.Printf("throwing the %s ball\n", color)
	balls <- color
}
