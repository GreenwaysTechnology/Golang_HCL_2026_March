package main

import (
	"fmt"
	"net/http"
)

// user resource
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Users Get")
	case http.MethodPost:
		fmt.Fprintln(w, "Users Post")
	case http.MethodPut:
		fmt.Fprintln(w, "Users Put")
	case http.MethodDelete:
		fmt.Fprintln(w, "Users Delete")
	}
}

// customer resource
func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Customer Get")
	case http.MethodPost:
		fmt.Fprintln(w, "Customer Post")
	case http.MethodPut:
		fmt.Fprintln(w, "Customer Put")
	case http.MethodDelete:
		fmt.Fprintln(w, "Customer Delete")
	}
}

// product resources
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Product Get")
	case http.MethodPost:
		fmt.Fprintln(w, "Product Post")
	case http.MethodPut:
		fmt.Fprintln(w, "Product Put")
	case http.MethodDelete:
		fmt.Fprintln(w, "Product Delete")
	}
}
func main() {
	//start the webserver
	http.HandleFunc("/user/", UserHandler)
	http.HandleFunc("/customer/", CustomerHandler)
	http.HandleFunc("/product/", ProductHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
