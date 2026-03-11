package main

import "fmt"

func main() {
	age := 18
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teen")
	case age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")

	}
}
