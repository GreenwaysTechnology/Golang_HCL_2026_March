package main

import "fmt"

func main() {
	users := make(map[string]int)
	//create
	users["Alice"] = 100
	users["Bob"] = 200
	//read
	fmt.Println(users["Alice"])
	fmt.Println(users["Bob"])
	//update
	users["Alice"] = 800
	fmt.Println(users["Alice"])
	//delete
	delete(users, "Alice")
	fmt.Println(users["Alice"])
}
