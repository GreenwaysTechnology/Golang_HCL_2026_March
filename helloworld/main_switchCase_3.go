package main

import "fmt"

func main() {
	switch num := 10; num {
	case 10:
		fmt.Println("10")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("default")
	}
}
