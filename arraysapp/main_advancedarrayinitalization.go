package main

import "fmt"

func main() {
	//size of the array is based on value which is decided during compile time.
	counters := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(counters)

	//index based initalzation
	arr := [5]int{
		0: 10,
		3: 90,
		4: 30,
	}
	fmt.Println(arr)
}
