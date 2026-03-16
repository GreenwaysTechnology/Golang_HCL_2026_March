package main

import "encoding/json"

func main() {
	response := struct {
		Status  string
		Message string
	}{
		Status:  "OK",
		Message: "User created successfully",
	}
	jsonData, _ := json.Marshal(response)
	println(string(jsonData))
}
