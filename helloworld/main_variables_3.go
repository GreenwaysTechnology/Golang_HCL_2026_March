package main

import "fmt"

func main() {
	fmt.Println("Variables Declaration")
	//type inference with shortcut : the type of variable is understood based on literal(value)
	name := "Subramanian Murugan"
	//var age uint16 = 46
	age := 46
	//var salary float64 = 500000
	salary := 500000.987
	status := true

	fmt.Println("name :", name)
	fmt.Println("age : ", age)
	fmt.Println("salary : ", salary)
	fmt.Println("status : ", status)
}
