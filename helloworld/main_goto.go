package main

import "fmt"

/*
*

		go syntax
		goto label
	    .....code.... This block of code is not executed
		label:
			....code....
*/
func main() {
	fmt.Println("start")
	goto END
	fmt.Println("this is goto")
END:
	fmt.Println("end")
}
