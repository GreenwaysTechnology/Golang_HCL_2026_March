package main

import "fmt"

func main() {
	ch := make(chan string, 2) //buffer size is 2
	ch <- "Hello"
	ch <- "World"
	ch <- "!"
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
