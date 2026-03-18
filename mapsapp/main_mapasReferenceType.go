package main

import "fmt"

func main() {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	fmt.Println(m1)
	m2 := m1
	fmt.Println(m2)
	m2["one"] = 1000
	fmt.Println(m1)
}
