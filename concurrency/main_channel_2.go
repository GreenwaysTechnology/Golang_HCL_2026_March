package main

import "fmt"

func sayHello(channel chan string) {
	channel <- "Hello ,How are you?"
}
func receiver(channel chan string) {
	msg := <-channel
	fmt.Printf("Message from channel: %s\n", msg)
}

func main() {
	channel := make(chan string)
	go sayHello(channel)
	receiver(channel)
}
