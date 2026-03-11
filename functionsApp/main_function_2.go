package main

import "fmt"

//function with parameters
/**
	func funcName(variable datatype){

   }
*/
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
func main() {
	greet("Subramanian Murugan")
}
