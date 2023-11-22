package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Total Items Packed:", PackItems(0))
}

func PackItems(totalItems int) int {
	const workers = 2
	const itemsPerWorker = 1000

	var wg sync.WaitGroup

	itemsPacked := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()
			// Simulate the worker packing items into boxes.
			for j := 0; j < itemsPerWorker; j++ {
				itemsPacked = totalItems
				// Simulate packing an item.
				itemsPacked++
				// Update the total items packed without proper synchronization.
				totalItems = itemsPacked
			}
		}(i)

	}

	// Wait for all workers to finish.
	wg.Wait()

	return totalItems
}
