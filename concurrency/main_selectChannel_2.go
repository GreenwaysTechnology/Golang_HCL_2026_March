package main

import (
	"fmt"
	"time"
)

func fetchFromCloud(serverName string, ch chan<- string, delay time.Duration) {
	time.Sleep(delay)
	ch <- "Data from" + serverName
}

func main() {
	serverA := make(chan string)
	serverB := make(chan string)
	go fetchFromCloud("http//www.proxy1.com", serverA, time.Millisecond*100)
	go fetchFromCloud("http//www.proxy2.com", serverB, time.Millisecond*50)
	//select statement acts as a race: first channel to send a value wins
	select {
	case messageOne := <-serverA:
		fmt.Println("SUCCESS", messageOne)
	case messageTwo := <-serverB:
		fmt.Println("SUCCESS", messageTwo)
	case <-time.After(time.Second * 1): //safety to avoid deadlocks
		fmt.Println("Timeout: Both servers were too slow")
	}
}
