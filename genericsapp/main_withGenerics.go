package main

import "fmt"

// generics function
// func FUnctionName[placeHolder type](args Type) Type {}
func sum[T int | float64 | int64](a, b T) T {
	return a + b
}
func main() {
	fmt.Println("Integer Addition", sum(1, 2))
	fmt.Println("Floating point Addition", sum(1.0, 2.78))
	//fmt.Println("String Values", sum("test", "test"))
}
