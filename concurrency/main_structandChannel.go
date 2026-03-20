package main

import "fmt"

type Message struct {
	ID   int
	Text string
}

func sendData(ch chan Message) {
	ch <- Message{1, "Hello"}
	ch <- Message{2, "World"}
	close(ch)
}
func main() {
	channel := make(chan Message)
	go sendData(channel)
	for message := range channel {
		fmt.Println(message.ID, message.Text)
	}
}
