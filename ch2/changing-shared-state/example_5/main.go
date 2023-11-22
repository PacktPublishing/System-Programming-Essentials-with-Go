package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Mutex{}
	fmt.Println("Total Items Packed:", PackItems(&m, 0))
}

func PackItems(m *sync.Mutex, totalItems int) int {
	const workers = 2
	const itemsPerWorker = 1000

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < itemsPerWorker; j++ {
				m.Lock()
				itemsPacked := totalItems
				itemsPacked++
				totalItems = itemsPacked
				m.Unlock()
			}
		}(i)
	}

	// Wait for all workers to finish.
	wg.Wait()

	return totalItems
}
