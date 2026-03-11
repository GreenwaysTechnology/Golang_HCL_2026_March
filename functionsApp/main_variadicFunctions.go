package main

import "fmt"

/*
*

	variable args - var args - variadic
*/
func sum(numbers ...int) int {
	total := 0
	//_ indicates index
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	fmt.Printf("%d\n", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
