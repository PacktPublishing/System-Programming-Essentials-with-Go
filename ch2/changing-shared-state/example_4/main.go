package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Total Items Packed:", PackItems(0))
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
		}(i)

	}

	// Wait for all workers to finish.
	wg.Wait()
	atomic.SwapInt32(&totalItems, itemsPacked)

	return totalItems
}
