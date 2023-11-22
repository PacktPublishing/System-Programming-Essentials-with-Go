package main

import (
	"fmt"
	"sync"
)

func main() {
	signalChannel := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 is waiting for a signal...")
		<-signalChannel
		fmt.Println("Goroutine 1 received the signal and is now doing something.")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 is about to send a signal.")
		signalChannel <- true
		fmt.Println("Goroutine 2 sent the signal.")
	}()

	wg.Wait()
	fmt.Println("Both goroutines have finished.")
}
