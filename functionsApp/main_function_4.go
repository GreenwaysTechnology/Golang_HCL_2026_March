package main

import "fmt"

/*
*
function with multiple values return

	func functionName(args) (type,type,type...){
		return var1,var2,var3,varN..
	}
*/
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient : %d  Remainder %d  \n", quotient, remainder)
}
