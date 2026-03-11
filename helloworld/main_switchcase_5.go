package main

import "fmt"

//fallthrough forces execution of next case

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("0")

	}
}
