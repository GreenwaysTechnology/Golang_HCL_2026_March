package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "!"
	}()
	//how to read muliple messages
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	
}
