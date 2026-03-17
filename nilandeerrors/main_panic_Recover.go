package main

import "fmt"

func safe() {
	fmt.Println("safe")
	//recover logic must be written inside anonmous defer function
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovered", e)
		}
	}()
	panic("something bad happened")
	//the code written after panic even though recovery not reachable
	fmt.Println("your task")
}
func doSomething() {
	fmt.Println("doSomething")
}

func main() {
	safe()
	doSomething()
}
