package main

import "fmt"

func main() {
	x := 10
	//pointer variable
	var p *int
	//assign the address of something(x)
	p = &x

	fmt.Println("Value of x:", x)
	fmt.Println("The address of X via Pointer Variable called p:", p)
	fmt.Println("The address of X :", &x)
	fmt.Println("Value of ptr (indirect value of x):", *p)
	fmt.Println("The value  of X via *(&x):", *(&x))
	//fmt.Println("Value of ptr (indirect value of x):", *x) //This is invalid
}
