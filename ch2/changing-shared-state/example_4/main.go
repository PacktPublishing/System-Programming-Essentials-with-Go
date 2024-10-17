package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	times := 0

	for {
		times++
		counter := PackItems(0)
		if counter != 2000 {
			log.Fatalf("it should be 2000 but found %d on execution %d", counter, times)
		}
	}
}

func PackItems(totalItems int32) int32 {
	const workers = 2
	const itemsPerWorker = 1000

	var wg sync.WaitGroup

	itemsPacked := int32(0)
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()
			// Simulate the worker packing items into boxes.
			for j := 0; j < itemsPerWorker; j++ {
				atomic.AddInt32(&itemsPacked, 1)
			}
			atomic.SwapInt32(&totalItems, itemsPacked)
		}(i)

	}

	// Wait for all workers to finish.
	wg.Wait()

	return totalItems
}
