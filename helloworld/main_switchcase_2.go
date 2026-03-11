package main

import "fmt"

func main() {
	day := 5
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Week days")
	default:
		fmt.Println("Weekend days")
	}
}
