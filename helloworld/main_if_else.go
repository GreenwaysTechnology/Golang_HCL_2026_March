package main

import "fmt"

func main() {
	age := 15
	if age >= 18 {
		fmt.Println("age>18")
	} else {
		fmt.Println("age<18")
	}
	num := 10
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
