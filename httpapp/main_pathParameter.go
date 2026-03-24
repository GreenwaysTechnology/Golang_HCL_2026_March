package main

import (
	"fmt"
	"net/http"
)

func main() {

	//old way of reading path parameters
	//http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
	//	///the prefix we are looking for
	//	prefix := "/user/"
	//	//ensure the path acutally starts with prefix
	//	if len(r.URL.Path) <= len(prefix) {
	//		http.Error(w, "User Id is Required (e.g.,/user/123", http.StatusBadRequest)
	//	}
	//	//slicing the string from the end of "/user/
	//	id := r.URL.Path[len(prefix):]
	//	id = strings.Split(id, "/")[0]
	//	fmt.Fprintln(w, "Customer:", id)
	//})

	//lastest way of reading path params go version 1.22+
	http.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "Customer:", id)
	})
	fmt.Println("Server has started....")
	http.ListenAndServe(":8080", nil)
}
