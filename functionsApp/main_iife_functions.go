package main

import "fmt"

/**
IIFE (immediately invoked function Expression)
*/

func main() {
	result := func(a, b int) int {
		return a + b
	}(10, 10) //invoking function
	fmt.Printf("result: %v \n", result)
	//without variable
	func() {
		fmt.Println("Hello im anonymous function")
	}()
}
