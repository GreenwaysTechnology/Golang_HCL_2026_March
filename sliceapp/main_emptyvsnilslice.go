package main

import "fmt"

// empty slice vs nil slice : empty and nil is not same
func main() {
	//nil slice
	var s []int
	fmt.Println(s)
	fmt.Println(s == nil)
	//Empty slice
	var s1 = []int{}
	fmt.Println(s1)
	fmt.Println(s1 == nil)
}
