package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	clownChannel := make(chan int, 3)
	clowns := 5

	// senders and receivers logic here!
	var wg sync.WaitGroup
	wg.Wait()
	fmt.Println("Circus car ride is over!")

	go func() {
		defer close(clownChannel)
		for clownID := range clownChannel {
			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Driver: Drove the car with %s inside\n", balloon)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Driver: Clown finished with %s, the car is ready for more!\n", balloon)
		}
	}()

	for clown := 1; clown <= clowns; clown++ {
		wg.Add(1)
		go func(clownID int) {
			defer wg.Done()
			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Clown %d: Hopped into the car with %s\n", clownID, balloon)
			select {
			case clownChannel <- clownID:
				fmt.Printf("Clown %d: Finished with %s\n", clownID, balloon)
			default:
				fmt.Printf("Clown %d: Oops, the car is full, can't fit %s!\n", clownID, balloon)
			}
		}(clown)
	}
}
