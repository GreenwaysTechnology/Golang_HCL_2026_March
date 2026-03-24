package main

import (
	"fmt"
	"net/http"
)

func main() {
	//router/ application logic
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	})
	http.HandleFunc("/hai", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hai")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})
	//start the webserver
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
