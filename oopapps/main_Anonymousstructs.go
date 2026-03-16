package main

import "fmt"

// anonmous struct return
/**
 func functionName() returnType {
	return value
}
func functionName () struct { name,age} {
	//function body
    return struct {variables}{intializedvalue}
}
*/
func getUser() struct {
	name string
	age  int
} { //function braces
	//create structus and initalize , return
	return struct {
		name string
		age  int
	}{name: "Subramanian", age: 46}
}

func main() {
	//anonymous structs
	person := struct {
		Name   string
		Age    int
		status bool
	}{
		Name:   "John Doe",
		Age:    42,
		status: true,
	}
	fmt.Println(person)
	user := getUser()
	fmt.Println(user.name, user.age)
}
