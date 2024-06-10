package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var g singleflight.Group
	var wg sync.WaitGroup

	fetchData := func(key string) (interface{}, error) {
		time.Sleep(2 * time.Second)
		return fmt.Sprintf("Data for key %s", key), nil
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result, err, shared := g.Do("my_key", func() (interface{}, error) {
				return fetchData("my_key")
			})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Goroutine %d got result: %v (shared: %v)\n", i, result, shared)
		}(i)
	}

	wg.Wait()
}
