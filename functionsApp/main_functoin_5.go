package main

import "fmt"

func divide_v2(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	quotient, remainder := divide_v2(10, 3)
	fmt.Printf("Quotient : %d  Remainder %d  \n", quotient, remainder)
}
