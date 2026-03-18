package main

import "fmt"

func main() {
	students := map[string]map[string]int{
		"Alice": {
			"math":            80,
			"science":         70,
			"computerscience": 90,
		},
		"Bob": {
			"math":            56,
			"science":         45,
			"computerscience": 40,
		},
	}
	fmt.Println(students)
	fmt.Println(students["Alice"]["math"])
}
