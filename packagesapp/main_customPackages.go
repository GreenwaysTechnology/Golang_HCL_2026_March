package main

//import module_name/packageName - this for custom packages and third party packages
import (
	"example.com/hcl/calculator"
	"example.com/hcl/util"
	"fmt"
)

func main() {
	fmt.Println(calculator.Myapp)
	result := calculator.Add(10, 10)
	fmt.Printf("Result Add : %d\n", result)
	result = calculator.Multiply(10, 10)
	fmt.Printf("Result Multiply : %d\n", result)
	result = calculator.Subtract(6, 10)
	fmt.Printf("Result Subtract : %d\n", result)
	result1 := calculator.Divide(45.9, 90.0)
	fmt.Printf("Result Divide : %f\n", result1)

	fmt.Println(util.GetInfo())
}
