package spewg

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestCacheInitialization(t *testing.T) {
	cache := NewCache(10)
	if cache == nil {
		t.Errorf("NewCache() = %v, want non-nil", cache)
	}
}

func TestCacheSetAndGetBehavior(t *testing.T) {
	cache := NewCache(10)
	cache.Set("key1", "value1", 1*time.Hour)

	value, found := cache.Get("key1")
	if !found || value != "value1" {
		t.Errorf("Get() = %v, %v, want %v, %v", value, found, "value1", true)
	}
}

func TestCacheGetNonExistentKeyBehavior(t *testing.T) {
	cache := NewCache(10)

	_, found := cache.Get("nonExistentKey")
	if found {
		t.Errorf("Get() = %v, want %v", found, false)
	}
}

func TestCacheSetOverwritesValueBehavior(t *testing.T) {
	cache := NewCache(10)
	cache.Set("key1", "value1", 1*time.Hour)
	cache.Set("key1", "value2", 1*time.Hour)

	value, _ := cache.Get("key1")
	if value != "value2" {
		t.Errorf("Get() = %v, want %v", value, "value2")
	}
}

func TestCacheSetUpdatesExpiryTime(t *testing.T) {
	cache := NewCache(2)
	cache.Set("key1", "value1", 1*time.Second)
	time.Sleep(2 * time.Second)
	_, found := cache.Get("key1")
	if found {
		t.Errorf("Get() = %v, want %v", found, false)
	}
	cache.Set("key1", "value1", 1*time.Hour)
	_, found = cache.Get("key1")
	if !found {
		t.Errorf("Get() = %v, want %v", found, true)
	}
}

func TestCacheEvictsLRU(t *testing.T) {
	cache := NewCache(2)
	cache.Set("key1", "value1", 1*time.Hour)
	cache.Set("key2", "value2", 1*time.Hour)
	cache.Set("key3", "value3", 1*time.Hour)
	_, found := cache.Get("key1")
	if found {
		t.Errorf("Get() = %v, want %v", found, false)
	}
}

func TestCacheEvictsExpiredItems(t *testing.T) {
	cache := NewCache(2)
	cache.Set("key1", "value1", 1*time.Second)
	time.Sleep(2 * time.Second)
	cache.evictExpiredItems()
	_, found := cache.Get("key1")
	if found {
		t.Errorf("Get() = %v, want %v", found, false)
	}
}

func TestCacheConcurrency(t *testing.T) {
	cache := NewCache(10) // Set a small capacity to induce eviction
	var wg sync.WaitGroup

	// Number of concurrent goroutines
	numGoroutines := 10

	// Number of operations per goroutine
	opsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < opsPerGoroutine; j++ {
				key := strconv.Itoa(j) // Use simple keys for testing
				value := "value" + key

				// Set with a short TTL to test eviction
				cache.Set(key, value, 1*time.Millisecond)
				time.Sleep(1 * time.Millisecond) // Add slight delay for TTL to expire

				// Get should either return the value or "", false (if expired/evicted)
				val, found := cache.Get(key)
				if found && val != value {
					t.Errorf("Unexpected value for key %s: got %s, want %s", key, val, value)
				}
			}
		}()
	}

	wg.Wait()
}
