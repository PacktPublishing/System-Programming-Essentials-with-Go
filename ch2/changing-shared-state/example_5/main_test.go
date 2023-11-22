package main

import (
	"sync"
	"testing"
)

func TestPackItems(t *testing.T) {
	m := sync.Mutex{}
	totalItems := PackItems(&m, 2000)
	expectedTotal := 2000
	if totalItems != expectedTotal {
		t.Errorf("Expected total: %d, Actual total: %d", expectedTotal, totalItems)
	}
}
