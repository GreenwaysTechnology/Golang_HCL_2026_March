package main

import "fmt"

/*
*

	closure is an anonymous function that captures/remembers variables from its surrounding
	scope

scopes

	-functions scope - any variables/functions/code within function -local scope
	-package scope - any variables/functions/code outside function within package
*/

// package scope
var x int = 10

func getX() {
	fmt.Printf(" x inside getX %d\n", x)
}
func main() {
	//local scope
	var x int = 20
	fmt.Printf("x inside Main : %d\n", x)
	getX()
}
