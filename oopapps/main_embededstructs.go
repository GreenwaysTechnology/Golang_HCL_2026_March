package main

import "fmt"

type Employee struct {
	Name string
	//Nested st
	Address struct {
		City  string
		State string
	}
}

func main() {
	var emp Employee
	emp.Name = "Jack"
	emp.Address.City = "New York"
	emp.Address.State = "New York"
	fmt.Println(emp.Name, emp.Address.City, emp.Address.State)
}
