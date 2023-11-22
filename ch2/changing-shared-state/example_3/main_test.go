package main

import (
	"testing"
)

func TestPackItems(t *testing.T) {
	totalItems := PackItems(2000)
	expectedTotal := 2000
	if totalItems != expectedTotal {
		t.Errorf("Expected total: %d, Actual total: %d", expectedTotal, totalItems)
	}
}
