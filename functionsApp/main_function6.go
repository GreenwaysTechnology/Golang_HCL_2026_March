package main

import "fmt"

/*
*
Named returns

	func Name() (variableName type) {
		return
	}
*/
func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	result := multiply(1, 2)
	fmt.Printf("Result : %d\n", result)
}
