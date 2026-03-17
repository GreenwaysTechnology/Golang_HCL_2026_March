package main

import "fmt"

func main() {
	//array length itself is consided as type
	//var nums1 [5]int
	//var nums2 [2]int
	//if nums1 == nums2 {
	//	fmt.Println("Arrays are equal")
	//} else {
	//	fmt.Println("Arrays are not equal")
	//}
	var nums1 [5]int
	var nums2 [5]int
	if nums1 == nums2 {
		fmt.Println("Arrays are equal")
	} else {
		fmt.Println("Arrays are not equal")
	}

}
