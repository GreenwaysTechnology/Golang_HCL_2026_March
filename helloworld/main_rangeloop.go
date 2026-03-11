package main

import "fmt"

//range loop with strings
/**
for index,value:=range collection/string/arrays/map
*/
func main() {
	str := "Go lang"
	for index, ch := range str {
		fmt.Println("index : ", index, "Value : ", ch)
	}
}
