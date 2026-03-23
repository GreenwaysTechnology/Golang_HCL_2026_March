package main

import "fmt"

//the output of one function will be input to another function
// output is chan , input is chan

// produces series of values - outbound channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq(in <-chan int) <-chan int {
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
	// Set up the pipeline.
	//c := gen(2, 3, 4, 5, 6, 10)
	//out := sq(c)
	////fmt.Println(<-out) // 4
	////fmt.Println(<-out) // 9
	//for value := range out {
	//	fmt.Println("value:", value)
	//}
	//short cut
	for n := range sq(gen(2, 3, 5, 6, 9, 10)) {
		fmt.Println(n) // 16 then 81
	}

}
