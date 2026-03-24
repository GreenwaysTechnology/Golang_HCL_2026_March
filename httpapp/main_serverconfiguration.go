package main

import "net/http"

func main() {
	//start the webserver
	server := http.Server{
		Addr: ":8080",
		
		Handler: nil,
	}
	//http.ListenAndServe(":8080", nil)
	server.ListenAndServe()
}
