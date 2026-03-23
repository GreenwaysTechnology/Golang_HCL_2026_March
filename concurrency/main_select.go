package main

import (
	"fmt"
	"time"
)

func main() {
	//create two channels
	ch1 := make(chan int)
	ch2 := make(chan error)
	//first go routine which sends number
	go func() {
		time.Sleep(20 * time.Second)
		ch1 <- 42
	}()
	go func() {
		time.Sleep(10 * time.Second)
		ch2 <- fmt.Errorf("Something bad happened")
	}()
	fmt.Println("Waiting for results and error")
	//select statement which waits for int and error channel who comes first that result
	//will be returned
	select {
	case value := <-ch1:
		fmt.Println("Results received", value)
	case err := <-ch2:
		fmt.Println("Error received", err)
	case <-time.After(time.Second * 5):
		fmt.Println("Timed out")
	}
}
