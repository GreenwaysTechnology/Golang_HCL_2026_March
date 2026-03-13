package main

import "fmt"

func main() {
	x := 10
	var p *int
	p = &x
	fmt.Println("Initial Value of x :", x)
	fmt.Println("The address of X via p(pointer) :", p)
	fmt.Println("The Value of X via pointer(p) :", *p)

	//reinitialization
	x = 20
	fmt.Println("After initialization Value of x :", x)
	fmt.Println("The address of X via p(pointer) :", p)
	fmt.Println("The Value of X via pointer(p) :", *p)

	//change the value of x via pointer : Dereferencing
	*p = 30
	fmt.Println("After deReferencing Value of x :", x)
	fmt.Println("The address of X via p(pointer) :", p)
	fmt.Println("The Value of X via pointer(p) :", *p)

}
