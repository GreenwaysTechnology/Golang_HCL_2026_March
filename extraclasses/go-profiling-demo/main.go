package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// This function simulates a CPU-heavy task with many allocations
func heavyTask() {
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024*1024) // Allocate 1MB
		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println("Application started...")
	heavyTask()
	fmt.Println("Application finished.")
}