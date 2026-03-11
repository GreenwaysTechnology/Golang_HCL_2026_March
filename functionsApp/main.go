package main

import "fmt"

/**
returning function from another function which is called  closure
*/

func counter() func() int {
	i := 0 //this is outer function variable
	//inner function : this is now closure
	return func() int {
		i++
		return i
	}
}

func main() {
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	nextState := counter()
	fmt.Println(nextState())
	fmt.Println(nextState())

}
