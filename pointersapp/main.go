package main

import "fmt"

// get the memory address using &
func getAddressUsingAnd() *int {
	x := 10
	//var x int
	return &x //This may leak data
}

// get the memory address using new
func getAddress() *int {
	//var x int
	x := 10
	return new(x) //this never leaks data
}
func main() {
	fmt.Println(getAddressUsingAnd(), *getAddressUsingAnd())
	fmt.Println(getAddressUsingAnd(), *getAddressUsingAnd())
}
