package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		if name == "" {
			name = "guest"
		}
		fmt.Fprintln(w, "Hello,", name)

	})
	fmt.Println("Server has started....")
	http.ListenAndServe(":8080", nil)
}
