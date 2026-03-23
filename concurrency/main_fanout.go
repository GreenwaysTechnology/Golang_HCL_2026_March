package main

import "fmt"

func gen_1(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq_1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			//find squre root of n and write into a channel
			out <- n * n
		}
		close(out)
	}()
	return out
}
func main() {
	//Fan out pattern: same channel input is distributed to multiple channel
	in := gen_1(2, 3)
	// FAN-OUT (2 workers)
	c1 := sq_1(in)
	c2 := sq_1(in)
	// Consume each separately
	for n1 := range c1 {
		fmt.Println("c1:", n1)
	}
	for n := range c2 {
		fmt.Println("c2:", n)
	}

}
