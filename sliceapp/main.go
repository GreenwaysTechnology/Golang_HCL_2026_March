package main

import "fmt"

func create() []int {
	s := []int{1, 2, 3}
	return s
}

func main() {
	//if you delcare and use slice within function that will be in stack
	s := []int{1, 2, 3}
	fmt.Println(s)
	//this is heap allocation beacuse returned slice must surive even after function
	//exits
	fmt.Println(create())

	//this slice may be allocated in heap
	slice := make([]int, 10000000)
	fmt.Println(len(slice))
}
