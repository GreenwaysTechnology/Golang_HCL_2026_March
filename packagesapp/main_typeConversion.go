package main

import (
	"fmt"
	"strconv"
)

/**
	Type conversion means converting a value from one data type to another data type
	Go does not support automatic(implicit) type conversion
	conversion syntax:
	 TargetType(value)
eg
	float64(10)
	int(3.14)
    string(65)
*/

func main() {
	//integer to float
	var a int = 10
	var b float64 = float64(a)
	fmt.Println("a=", a, "b=", b)
	//float to integer
	var x float64 = 9.87
	var y = int(x)
	fmt.Println("x=", x, "y=", y)

	//number to string conversion using rune
	num := 65
	str := string(num)
	fmt.Println("num=", num, "str=", str)

	//string to number
	cost := "145"
	//int cant be used to convert string to number
	//strconv
	var c, _ = strconv.Atoi(cost)
	//strconv.ParseInt(cost, 10, 10)
	fmt.Println("cost=", cost, "converted Cost=", c)
}
