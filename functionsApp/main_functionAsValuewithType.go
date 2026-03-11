package main

import "fmt"

func main() {
	//with type
	//add is variable
	//func(int, int) int => is type
	//func(a, b int) int {
	//		return a + b
	//	} => is value
	//var add func(int, int) int = func(a, b int) int {
	//	return a + b
	//}
	var add func(int, int) int
	fmt.Println(add)
	add = func(a, b int) int {
		return a + b
	}
	fmt.Printf("result is %d\n", add(1, 2))

	var subtract = func(a, b int) int {
		return a - b
	}
	fmt.Printf("result is %d\n", subtract(10, 2))

}
