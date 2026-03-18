package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := map[int]User{
		1: {Name: "Alice", Age: 30},
		2: {Name: "Bob", Age: 40},
		3: {Name: "Charlie", Age: 60},
		4: {Name: "David", Age: 60},
		5: {Name: "John", Age: 60},
	}
	for _, v := range users {
		fmt.Println("Name: ", v.Name, "Age: ", v.Age)
	}
}
