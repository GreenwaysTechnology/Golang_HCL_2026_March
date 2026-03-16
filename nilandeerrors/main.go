package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}
func login(username, password string) (string, error) {
	if username != "admin" && password != "admin" {
		return "", errors.New("login failed")
	}
	return "login success", nil
}

func main() {
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	//status, errorCode := login("admin", "admin")
	status, errorCode := login("foo", "bar")
	if errorCode != nil {
		fmt.Println(errorCode)
	}
	fmt.Println(status)

}
