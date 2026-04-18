package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Initialize the concurrent map
	var m sync.Map

	// 2. Store data (Key, Value)
	// These can be any type (interface{})
	m.Store("user_1", "Alice")
	m.Store("user_2", "Bob")

	// 3. Load data
	val, ok := m.Load("user_1")
	if ok {
		fmt.Println("Found:", val)
	}

	// 4. LoadOrStore: Atomic "Get or Create"
	// Returns actual value and true if loaded, false if stored
	actual, loaded := m.LoadOrStore("user_3", "Charlie")
	fmt.Printf("Loaded: %v, Value: %v\n", loaded, actual)

	// 5. Delete data
	m.Delete("user_2")

	// 6. Range: Iterate over the map
	// The function must return true to keep iterating
	fmt.Println("Current Map Contents:")
	m.Range(func(key, value any) bool {
		fmt.Printf(" - %s: %s\n", key, value)
		return true
	})
}