package main

import "fmt"

func main() {
	p := &struct {
		Name string
		Age  int
	}{
		Name: "John Doe",
		Age:  18,
	}
	fmt.Println(p.Name, p.Age)
}
