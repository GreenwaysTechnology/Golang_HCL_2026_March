package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Incrementing....")
	atomic.AddInt32(&counter, 1) //atomic operation
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Value %d: Finished!\n", counter)
}
