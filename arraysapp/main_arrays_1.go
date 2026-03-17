package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//declared array
	var nums [5]int
	fmt.Println("Integer Array : ", nums)
	var prices [5]float64
	fmt.Println("Float array : ", prices)

	var names [5]string
	fmt.Println("String Array : ", names)
	var persons [5]Person
	fmt.Println("Person Array : ", persons)
}
