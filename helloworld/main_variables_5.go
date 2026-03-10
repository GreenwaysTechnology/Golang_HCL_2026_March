package main

import "fmt"

func main() {
	//with type
	//var (
	//	name   string  = "Subramanian Murugan"
	//	age    int     = 46
	//	salary float64 = 500000.987
	//	status bool    = true
	//)
	//with type inference
	var (
		name   = "Subramanian Murugan"
		age    = 46
		salary = 500000.987
		status = true
	)
	fmt.Println("name :", name)
	fmt.Println("age :", age)
	fmt.Println("salary :", salary)
	fmt.Println("status :", status)
}
