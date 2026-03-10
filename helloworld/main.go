package main

import "fmt"

func main() {
	score := 99
	if score >= 90 {
		fmt.Println("Grade A!")
	} else if score >= 75 {
		fmt.Println("Grade B!")
	} else {
		fmt.Println("Grade C!")
	}

}
