package main

import (
	"fmt"
	"sync"
)

func gen_3(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq_3(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// FAN-IN: merge channels
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := gen_3(2, 3, 4, 5)

	// FAN-OUT
	c1 := sq_3(in)
	c2 := sq_3(in)

	// FAN-IN
	out := merge(c1, c2)

	// Single consumer
	for n := range out {
		fmt.Println("merged:", n)
	}

}
