package main

import "fmt"

/**
Anonymous function is a function without name
*/

func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("%d\n", add(1, 2))
}
