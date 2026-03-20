package main

import (
	"concurrency/goroutineone"
	"fmt"
	"sync"
)

func main() {
	//A WaitGroup is a counting semaphore typically used to wait for a group of
	//goroutines or tasks to finish.
	var wg sync.WaitGroup

	counter := 100
	for i := 0; i <= counter; i++ {
		wg.Add(1) // +1 =1+1+1+1....
		go goroutineone.HelloNew("Hello_", &wg)
	}
	fmt.Println("Main go routine")
	//pause the main or any goroutine to finish
	wg.Wait()
	fmt.Println("Program finished")
}
