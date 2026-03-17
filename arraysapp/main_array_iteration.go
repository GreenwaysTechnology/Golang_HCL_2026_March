package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	//classical for loop
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	//range loop
	//with index useage
	for index, value := range array {
		fmt.Println(index, value)
	}
	//_ unused variable
	for _, v := range array {
		fmt.Println(v)
	}
}
