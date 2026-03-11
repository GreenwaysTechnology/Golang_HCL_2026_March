package main

import "fmt"

/*
*

	closure is an anonymous function that captures/remembers variables from its surrounding
	scope

scopes

	-functions scope - any variables/functions/code within function -local scope
	-package scope - any variables/functions/code outside function within package
*/

func main() {
	//closure
	x := 10 //the variable is declared inside main - outter function
	//inner function
	innerFunc := func() {
		//now inner function accessing outter function variable, this is called
		//closure
		fmt.Printf("x inner func : %d\n", x)
	}
	innerFunc()
}
