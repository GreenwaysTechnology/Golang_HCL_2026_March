package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. Create a base context and a cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// 2. Start the long-running Goroutine
	go longRunningTask(ctx)

	// 3. Let it run for a bit
	time.Sleep(3 * time.Second)

	// 4. Trigger the cancellation
	fmt.Println("Main: Stopping the goroutine now...")
	cancel()

	// Give it a moment to cleanup before the program exits
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting.")
}

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 5. Receive the cancellation signal
			fmt.Println("Task: Received cancel signal! Cleaning up...")
			return
		default:
			// Simulate work
			fmt.Println("Task: Working...")
			time.Sleep(1 * time.Second)
		}
	}
}