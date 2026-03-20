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

	//How many goroutines are there /are going to be lanuched

	//wg.Add(2)
	go goroutineone.HelloNew("how are you", &wg)
	go goroutineone.HelloNew("i am fine", &wg)

	fmt.Println("Main go routine")
	//pause the main or any goroutine to finish
	wg.Wait()
	fmt.Println("Program finished")
}
