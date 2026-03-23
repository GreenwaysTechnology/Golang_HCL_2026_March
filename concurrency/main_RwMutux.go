package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

// read
func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock() // Multiple readers OK
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// write
func (c *Cache) Set(key string, val int) {
	c.mu.Lock() // Exclusive writer
	defer c.mu.Unlock()
	c.data[key] = val
}

func main() {
	cache := &Cache{data: make(map[string]int)}
	// Simulate read-heavy workload (90% reads)
	var wg sync.WaitGroup
	// 90 reader goroutines
	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				cache.Get("key") // Fast - multiple readers
			}
		}(i)
	}
	// 10 writer goroutines (rare)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond) // Simulate slow writes
			cache.Set("key", id)
		}(i)
	}
	start := time.Now()
	wg.Wait()
	fmt.Printf("RWMutex: %d readers + %d writers took %v\n", 90, 10, time.Since(start))

}
