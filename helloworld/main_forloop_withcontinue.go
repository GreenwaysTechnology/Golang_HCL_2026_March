package main

import "fmt"

// continue
func main() {
	for i := 1; i <= 5; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
