package main

import "fmt"

//comparable Constraint is built in interface - for comparing generic values using ==
//=!

func Equal[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	result := Equal[int](1, 1)
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	result = Equal[string]("hello", "world")
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	result = Equal[float32](10.90, 10.90)
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
