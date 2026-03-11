package main

import "fmt"

func login(resolve func(status string) int, reject func(err string) int) {
	loggedIn := true
	var statuscode int
	if loggedIn {
		statuscode = resolve("Login success")
	} else {
		statuscode = reject("Login failed")
	}
	fmt.Println(statuscode)
}

func main() {
	//passed two function as parameter
	login(func(status string) int {
		//fmt.Printf("Login Success")
		fmt.Println(status)
		return 200
	}, func(err string) int {
		//fmt.Printf("Login Failed")
		fmt.Println(err)
		return 500
	})
}
