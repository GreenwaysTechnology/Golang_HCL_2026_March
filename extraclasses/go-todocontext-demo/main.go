package main

import (
	"context"
	"fmt"
	"time"
)

// --- DATABASE LAYER (Modern/Updated) ---
// This function is "Correct" - it takes a context and respects it.
func SaveOrderToDB(ctx context.Context, orderID string) error {
	select {
	case <-time.After(500 * time.Millisecond): // Simulate DB write
		fmt.Printf("Successfully saved order %s to database\n", orderID)
		return nil
	case <-ctx.Done(): // If the context is cancelled or times out
		return ctx.Err()
	}
}

// --- BUSINESS LOGIC LAYER (In Progress) ---
// TODO: This function should be updated to accept (ctx context.Context, id string)
// But for now, it doesn't, so we use context.TODO() inside it.
func ProcessOrder(orderID string) {
	fmt.Printf("Processing order: %s...\n", orderID)

	// We NEED a context to call SaveOrderToDB.
	// Since ProcessOrder doesn't have one passed in yet, we use TODO.
	ctx := context.TODO()

	err := SaveOrderToDB(ctx, orderID)
	if err != nil {
		fmt.Printf("Error processing order: %v\n", err)
		return
	}

	fmt.Println("Order flow complete.")
}

// --- MAIN ENTRY POINT ---
func main() {
	// Let's process an order.
	// Notice that main doesn't have a context to give ProcessOrder yet.
	ProcessOrder("ORD-12345")
}
