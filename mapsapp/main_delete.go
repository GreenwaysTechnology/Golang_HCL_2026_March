package main

import "fmt"

func main() {
	m := map[string]int{"1": 1, "2": 2}
	fmt.Println(m)
	delete(m, "1")
	fmt.Println(m)
}
