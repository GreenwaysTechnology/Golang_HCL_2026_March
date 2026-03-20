package main

import "fmt"

func publisher(channel chan string) {
	//channel <- "Hello ,How are you?"
}
func subscriber(channel chan string) {
	msg := <-channel
	fmt.Printf("Message from channel: %s\n", msg)
}

func main() {
	channel := make(chan string)
	go publisher(channel)
	subscriber(channel)
}
