package main

import "fmt"

// Custom Type
type Salary float32
type Celcius float64

func main() {
	//var employeeSalary float32 = 10000
	var employeeSalary Salary = 10000
	var wage Salary = 1000
	fmt.Println(employeeSalary, wage)
	var temp Celcius = 100
	println(temp)
}
