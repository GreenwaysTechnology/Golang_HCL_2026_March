package main

import (
	"fmt"
	"sync"
)

// object which holds data
type SafeCounter struct {
	mu    sync.Mutex
	count int //data to be protected from multiple go routines access
}

// function to return counter
func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

// function to increment counter
func (sc *SafeCounter) Inc() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}
func worker(wg *sync.WaitGroup, counter *SafeCounter) {
	defer wg.Done()
	counter.Inc()
}
func main() {
	//create SafeCounter
	counter := SafeCounter{
		count: 0,
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &counter)
	}
	wg.Wait()
	fmt.Printf("%d\n", counter.Value())
}
