package main

import "fmt"

//package level variables

func main() {
	fmt.Println("VariablesDeclaration")
	//function level variables
	var name string = "Subramanian Murugan"
	var age uint16 = 46
	var salary float64 = 500000
	var status bool = true

	fmt.Println("name :", name)
	fmt.Println("age : ", age)
	fmt.Println("salary : ", salary)
	fmt.Println("status : ", status)
}
