package main

import "fmt"

// without generics
func sumInt(a, b int) int {
	return a + b
}
func sumFloat(a, b float64) float64 {
	return a + b
}

func main() {
	fmt.Println("Integer Addition", sumInt(1, 2))
	fmt.Println("Floating point Addition", sumFloat(1.0, 2.78))
}
