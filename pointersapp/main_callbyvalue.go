package main

import "fmt"

func squareVal(v int) {
	v *= v
	fmt.Println("The address of v ", &v, "The value of V", v)
}

func main() {
	a := 4
	//copy of the value : by pass value
	fmt.Println("Before squareVal:", a)
	squareVal(a)
	fmt.Println("The address of a", &a, "The value of a", a)
	fmt.Println("After squareVal:", a)
}
