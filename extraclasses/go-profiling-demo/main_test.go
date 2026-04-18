package main

import "testing"

func BenchmarkHeavyTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heavyTask()
	}
}