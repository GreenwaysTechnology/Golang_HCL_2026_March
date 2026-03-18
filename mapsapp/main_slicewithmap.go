package main

import "fmt"

func main() {
	m := map[string][]string{
		"eng":      {"Math", "physics", "chemistry", "Major"},
		"arts":     {"language", "history", "science"},
		"medicine": {"botany", "biology", "chemistry"},
	}
	for key, v := range m {
		fmt.Println("Subject:", key)
		for _, value := range v {
			fmt.Println(value)
		}
		fmt.Println(".......................")

	}
}
