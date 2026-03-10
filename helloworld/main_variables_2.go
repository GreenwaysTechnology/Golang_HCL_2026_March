package main

import "fmt"

func main() {
	fmt.Println("VariablesDeclaration")
	//type inference : the type of variable is understood based on literal(value)
	var name = "Subramanian Murugan"
	//var age uint16 = 46
	var age = 46
	//var salary float64 = 500000
	var salary = 500000.987
	var status = true

	fmt.Println("name :", name)
	fmt.Println("age : ", age)
	fmt.Println("salary : ", salary)
	fmt.Println("status : ", status)
}
