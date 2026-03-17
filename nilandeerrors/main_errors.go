package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		//return 0, errors.New("divide by zero")
		return 0, fmt.Errorf("the Parameter %d cant be divided by 0", a)
	}
	return a / b, nil
}
func login(username, password string) (string, error) {
	if username != "admin" && password != "admin" {
		//return "", errors.New("login failed")
		return "", fmt.Errorf("username or password error")
	}
	return "login success", nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	//status, errorCode := login("admin", "admin")
	status, errorCode := login("foo", "bar")
	if errorCode != nil {
		fmt.Println(errorCode)
		return
	}
	fmt.Println(status)

}
