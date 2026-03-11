package main

import "fmt"

/*
*
func funcName(variable datatype) datatype {

}
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	result := add(1, 2)
	fmt.Printf("Result: %d\n", result)
}
