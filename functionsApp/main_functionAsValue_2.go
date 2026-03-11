package main

import "fmt"

// variable with inference
var hai = func() {
	fmt.Println("hai")
}

func main() {
	hai()
}
