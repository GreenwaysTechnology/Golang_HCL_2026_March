package main

import "fmt"

// used to delay the execution of a function until the surrending function returns
func main() {
	fmt.Println("start")
	defer fmt.Println("Delayed function")
	fmt.Println("end")
}
