package main

import (
	"fmt"
	"time"
)

func Hello(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	//lanuch two goroutines, it runs independently
	go Hello("world")
	fmt.Println("Main go routine")
	go Hello("Hello")
	//block or wait the main go routine, otherwise main exits before Hello go routine
	//prints
	time.Sleep(1000 * time.Millisecond)

}
