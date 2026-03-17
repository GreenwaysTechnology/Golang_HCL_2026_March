package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 90
	fmt.Println(a)
	fmt.Println(b)

}
