package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	s := arr[2:5]
	fmt.Println(s)

	//slice declaration: slice has no size mentioned
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//using make api
	//
	var length int = 10
	var capacity int = 20
	sliceMake := make([]int, length, capacity)
	fmt.Println("capacity:", cap(sliceMake))
	fmt.Println("length:", len(sliceMake))
	fmt.Println("slice:", slice)
}
