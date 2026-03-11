package main

import "fmt"

func main() {
	//how closure can read and modify the outter scope variable
	counter := 0
	//closure
	increment := func() {
		counter++
		fmt.Println("counter =", counter)
	}
	increment()
	increment()
	increment()
	increment()
}
