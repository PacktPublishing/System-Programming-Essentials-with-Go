package main

import (
	"context"
	"testing"
	"time"

	"github.com/alexrios/timer/v2"
)

func TestStopwatch(t *testing.T) {
	sw := &timer.Stopwatch{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sw.Start(ctx); err != nil {
		t.Fatalf("Failed to start stopwatch: %v", err)
	}

	time.Sleep(1 * time.Second)
	sw.Stop()

	elapsed, err := sw.Elapsed()
	if err != nil {
		t.Fatalf("Failed to get elapsed time: %v", err)
	}
	if elapsed < 1*time.Second || elapsed > 2*time.Second {
		t.Errorf("Expected elapsed time around 1 second, got %v", elapsed)
	}
}
