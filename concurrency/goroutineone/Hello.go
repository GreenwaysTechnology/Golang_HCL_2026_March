package goroutineone

import (
	"fmt"
	"sync"
)

func HelloNew(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	//signal completion: meaning that this go routine finished work
	defer wg.Done()
}
