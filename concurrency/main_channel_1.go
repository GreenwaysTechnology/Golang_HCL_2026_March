package main

import "fmt"

func main() {
	fmt.Println("Channel Example")
	channel := make(chan string)
	//sender logic
	go func() {
		//write data called Hello into channel
		channel <- "Hello"
		fmt.Println("Producer has published data")
	}()
	//Receiver
	msg := <-channel
	fmt.Printf("Message from channel: %s\n", msg)

}
