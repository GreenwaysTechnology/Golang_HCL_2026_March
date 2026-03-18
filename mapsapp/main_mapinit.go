package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	fmt.Println(m == nil)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	fmt.Println(m)
}
