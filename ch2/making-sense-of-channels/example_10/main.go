package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	clownChannel := make(chan int, 3)
	driverChannel := make(chan struct{})
	clowns := 10
	circusRoundTrip := time.Second * 2 // How long the car takes to go around the circus

	fmt.Println("Announcer: The circus car is ready for clowns!")
	// The Driver is waiting for clowns to drive around.
	// They will wait until there are no more clowns. In other words, clownChannel is closed.
	go func() {
		// Receiving balls
		for ball := range clownChannel {
			balloon := fmt.Sprintf("Balloon %d", ball)
			fmt.Printf("Driver: Drove the car with %s inside\n", balloon)
			time.Sleep(circusRoundTrip)
			fmt.Printf("Driver: Clown finished with %s, the car is ready for more!\n", balloon)
		}

		fmt.Println("Driver: There is no more clowns to drive around! I'm done for the day!")
		driverChannel <- struct{}{}
	}()

	// Clowns are trying to get into the car.
	// They will keep trying in their own goroutine until they get in.
	wg := sync.WaitGroup{}
	wg.Add(clowns)
	for clown := 1; clown <= clowns; clown++ {
		go func(clownID int) {
			defer wg.Done()

			balloon := fmt.Sprintf("Balloon %d", clownID)
			for {
				select {
				case clownChannel <- clownID:
					fmt.Printf("Clown %d: Hopping into the car with %s\n", clownID, balloon)
					return
				default:
					fmt.Printf("Clown %d: Oops, the car is full, can't fit %s!\n", clownID, balloon)
				}
				retry()
			}
		}(clown)
	}

	// Announcer is waiting for all clowns start battling for a car spot.
	// They will wait until all clowns are done trying to get in the car.
	// In other words, the WaitGroup is done.
	// Once all clowns are done, the clownChannel is closed.
	go func() {
		fmt.Println("Announcer: All clowns are battling for the car!")
		wg.Wait()
		close(clownChannel)
	}()

	// Wait for the driver to finish driving all clowns around.
	<-driverChannel
	fmt.Println("Circus car ride is over!")
}

// retry is a function that simulates the clowns trying to get in the car chaotically
// by sleeping for a random amount of time between 2000 and 10000 milliseconds
func retry() {
	time.Sleep(time.Duration(max(rand.IntN(10000), 2000)) * time.Millisecond)
}
