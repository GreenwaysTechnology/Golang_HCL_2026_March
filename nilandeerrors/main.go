package main

import "fmt"

func A() {
	fmt.Println("A Start")
	B()
	fmt.Println("A End")
}
func B() {
	fmt.Println("B Start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered", err)
		}
	}()
	C()
	fmt.Println("B End")
}
func C() {
	//defer println("defered function in C")

	panic("something bad happened")
}
func main() {
	A()
	fmt.Println("continue")
}
