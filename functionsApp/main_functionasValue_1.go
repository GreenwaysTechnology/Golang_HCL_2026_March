package main

import "fmt"

/**
 functions are first class citizens.
 	functions are values like ints,floats,strings,booleans,objects
for eg:
 var a int =10
 var is keyworkd
 a is variable
 int is datatype
 10 is value/literal
 like 10 we can assign functions into variables.
 var myfunc = function defintion(){}
 we can invoke function using myfunc variable
*/

func main() {
	//with type
	var greet func() = func() {
		fmt.Println("Greeting")
	}
	greet()
	//without type : type inference
	var hello = func() {
		fmt.Println("Hello World")
	}
	hello()
	//without var and type: shortcut
	hai := func() {
		fmt.Println("Hi")
	}
	hai()
}
