package main

import "fmt"

func main() {
	//with type
	var x, y, z int = 1, 2, 3
	fmt.Println("x , y, Z", x, y, z)
	//with type inference
	var a, b, c = 1, 2, 3
	fmt.Println("a , b , c", a, b, c)
	var name, age, status = "subramanian", 46, true
	fmt.Println("name age, status", name, age, status)

}
