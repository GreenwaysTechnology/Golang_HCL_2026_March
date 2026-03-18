package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m["a"])
	m["a"] = 100
}
