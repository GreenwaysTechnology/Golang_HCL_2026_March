package main

import "fmt"

//go does not have while loop, so we can use for as while
/**
    declaration
	for condition {
		//do something
		//post/pre operations
    }
*/
func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
