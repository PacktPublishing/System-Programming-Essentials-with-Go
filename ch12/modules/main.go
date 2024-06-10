package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alexrios/timer/v2"
)

func main() {
	sw := &timer.Stopwatch{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sw.Start(ctx); err != nil {
		fmt.Printf("Failed to start stopwatch: %v\n", err)
		return
	}

	time.Sleep(1 * time.Second)
	sw.Stop()

	elapsed, err := sw.Elapsed()
	if err != nil {
		fmt.Printf("Failed to get elapsed time: %v\n", err)
		return
	}
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
