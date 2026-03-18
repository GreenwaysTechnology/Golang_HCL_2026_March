package main

import "fmt"

func main() {
	//literal initalization
	m := map[string]int{"1": 1, "2": 2}
	fmt.Println(m)
	//print map values
	fmt.Println(m["1"])
	//to check key and value
	value, ok := m["10"]
	if ok {
		fmt.Println("value is ", value)
	} else {
		fmt.Println("Key not found")
	}
}
