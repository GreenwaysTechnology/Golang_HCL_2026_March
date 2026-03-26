package main

import "fmt"

func PrintValue[T any](value T) {
	fmt.Println(value)
}
func main() {
	PrintValue("Hello World")
	PrintValue(10)
	PrintValue(98.8)
	PrintValue(false)
}
