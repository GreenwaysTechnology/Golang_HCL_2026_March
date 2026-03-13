package main

import "fmt"

//Escape Analysis : how to allocate memory in stack and heap

// stack allocation
func add(a, b int) int {
	return a + b
}
func test() *int {
	num := 10
	//return &num
	//num escapes to heap
	return &num
}

func main() {
	fmt.Println(add(1, 2))
	var p *int
	p = test()
	fmt.Println(p)
	fmt.Println(*p)
	*p = 90
	fmt.Println(*p)
}
