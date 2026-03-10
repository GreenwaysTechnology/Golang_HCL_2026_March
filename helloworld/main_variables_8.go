package main

import "fmt"

// varibles outside function : package level variables
// var age int = 18
// with type inference
///var age = 18
//short noatation: is not allowed outside functions
//age := 19

var (
	status = true
	age    = 18
)

func main() {
	//variables inside function are local variables
	var name = "Subramanian Murugan"
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("status:", status)
}
