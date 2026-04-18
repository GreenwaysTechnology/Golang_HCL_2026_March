package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	// We start 1000 goroutines to increment the same variable
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // This is the race condition!
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}