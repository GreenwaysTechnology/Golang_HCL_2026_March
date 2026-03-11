package main

import "fmt"

// Defered calls execute in Last-In-Firt-Out(stack) order
func main() {
	fmt.Println("START")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	defer fmt.Println("Fourth defer")
	fmt.Println("END")

}
