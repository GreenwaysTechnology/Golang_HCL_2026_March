package main

import "fmt"

func squareAdd(p *int) *int {
	//*p = *p * *p
	//dereferencing the a's value
	*p *= *p
	//fmt.Println("The value of p ", *p)
	return p // address of p // now p's value is allocated inside heap memory
}
func main() {
	a := 4
	fmt.Println("Before squareAdd a value:", a)
	var res = squareAdd(&a)
	fmt.Println("After squareAdd a value :", a)
	fmt.Println(res, *res)
}
