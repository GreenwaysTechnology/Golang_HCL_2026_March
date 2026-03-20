package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "!"
		//after publishing all messages, finally we have to close the channel
		close(channel)
	}()
	//how to read muliple messages using for..range
	for value := range channel {
		fmt.Println(value)
	}

}
