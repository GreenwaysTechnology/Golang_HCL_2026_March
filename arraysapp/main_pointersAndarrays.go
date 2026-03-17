package main

import "fmt"

func modify(array [3]int) {
	array[1] = 100
}
func modifyByPointer(array *[3]int) {
	array[1] = 100
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	b := a
	fmt.Println(b)
	b[0] = 100
	fmt.Println(b)
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
	array := [3]int{1, 2, 3}
	fmt.Println("Before Modify : ", array)
	modifyByPointer(&array)
	fmt.Println("After Modify : ", array)
}
