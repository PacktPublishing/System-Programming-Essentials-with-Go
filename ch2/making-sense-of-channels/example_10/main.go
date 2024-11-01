package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	clownChannel := make(chan int, 3)
	clowns := 5

	var wgClowns sync.WaitGroup
	var wgDriver sync.WaitGroup

	// sender/driver logic
	go func() {
		wgDriver.Add(1)
		defer wgDriver.Done()
		for clownID := range clownChannel {
			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Driver: Drove the car with %s inside\n", balloon)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Driver: Clown finished with %s, the car is ready for more!\n", balloon)
		}

		fmt.Println("Driver: I'm done for the day!")
	}()

	// receiver/clowns logic
	for clown := 1; clown <= clowns; clown++ {
		wgClowns.Add(1)
		go func(clownID int) {
			defer wgClowns.Done()
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

	wgClowns.Wait()
	close(clownChannel)
	wgDriver.Wait()
	fmt.Println("Circus car ride is over!")
}
