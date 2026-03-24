package main

import (
	"fmt"
	"net/http"
)

//func HomePagehandler(response http.ResponseWriter, request *http.Request) {
//	//response.Write([]byte("Welcome to the HomePage!"))
//	//fmt.Fprintf(response, "<h1>Hello World</h1>")
//	fmt.Fprintln(response, "Welcome to the HomePage!")
//}

func main() {
	//router/ application logic
	//http.HandleFunc("/", HomePagehandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	//start the webserver
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
