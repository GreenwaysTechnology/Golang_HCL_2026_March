package main

import "fmt"

func PrintPair[K any, V any](key K, value V) {
	fmt.Println(key, value)
}

func main() {
	PrintPair("one", 1)
	PrintPair(1, "Subramanian")
}
