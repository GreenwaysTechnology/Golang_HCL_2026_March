package main

import "fmt"

/**
Higher order function or HOF - passing function as parameter to another function
how? because is value
*/

func welcome(callback func()) {
	//invoke the function by invoking variable called callback
	callback()
}

// the function to be passed
func sayHello() {
	fmt.Println("Hello")
}

//	var h func() = func() {
//		fmt.Println("function with variables and types")
//	}
var h = func() {
	fmt.Println("function with variables and types")
}

func main() {
	welcome(sayHello)
	welcome(h)
	//anonymous function passing
	welcome(func() {
		fmt.Println("anonymous function")
	})
}
