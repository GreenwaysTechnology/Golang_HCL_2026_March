package main

import "fmt"

func main() {
	lang := "go"
	switch lang {
	case "go":
		fmt.Println("Go")
	case "python":
		fmt.Println("Python")
	case "java":
		fmt.Println("Java")
	case "c":
		fmt.Println("C")
	case "cpp":
		fmt.Println("C++")
	default:
		fmt.Println("Unknown language")

	}
}
