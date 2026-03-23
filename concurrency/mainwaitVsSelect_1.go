package main

import (
	"fmt"
	"sync"
	"time"
)

func processImage(id int, wg *sync.WaitGroup) {
	// Schedule the call to Done to ensure the counter decrements
	// even if the function panics.
	defer wg.Done()

	fmt.Printf("Image %d: Starting resize...\n", id)
	time.Sleep(time.Millisecond * 800) // Simulate processing time
	fmt.Printf("Image %d: Finished!\n", id)
}

func main() {
	var wg sync.WaitGroup
	imagesToProcess := 5

	for i := 1; i <= imagesToProcess; i++ {
		wg.Add(1) // Tell the counter we are starting 1 task
		go processImage(i, &wg)
	}

	fmt.Println("Main: Waiting for all images to finish...")
	wg.Wait() // This blocks until the counter hits 0
	fmt.Println("Main: All images processed. Ready to upload!")
}
