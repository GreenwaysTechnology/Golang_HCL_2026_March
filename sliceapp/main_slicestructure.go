package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice1 := array[1:4]
	slice2 := array[2:5]
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("Underlaying array", array)
	slice1[1] = 200
	fmt.Println("slice1", slice1)
	fmt.Println("Underlaying array", array)

	slice1[2] = 400
	fmt.Println("slice2", slice2)
	fmt.Println("Underlaying array", array)
	fmt.Println("slice1", slice1)

}
