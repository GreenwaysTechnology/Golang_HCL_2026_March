package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Fibonacci computes the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciHandler handles HTTP requests for Fibonacci calculations
func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("n")
	n, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%d", Fibonacci(n))
}