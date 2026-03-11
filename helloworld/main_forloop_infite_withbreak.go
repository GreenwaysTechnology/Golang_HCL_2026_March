package main

import "fmt"

func main() {
	i := 1
	//infinite loop : no condition is mentioned
	for {
		fmt.Println(i)
		i++
		if i == 100 {
			break
		}
	}
}
