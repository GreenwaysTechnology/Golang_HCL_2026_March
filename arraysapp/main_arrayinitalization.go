package main

import "fmt"

type Customer struct {
	Name string
	Age  int
}

func main() {
	//declared array
	//var nums [5]int = [5]int{1, 2, 3, 4, 5}
	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Integer Array : ", nums)

	var prices = [5]float64{10.9, 89.90, 90.87, 80.80, 50.0}
	fmt.Println("Float array : ", prices)

	var names = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("String Array : ", names)

	var persons = [5]Customer{{Name: "A", Age: 20}, {Name: "B", Age: 30}}
	fmt.Println("Person Array : ", persons)
}
