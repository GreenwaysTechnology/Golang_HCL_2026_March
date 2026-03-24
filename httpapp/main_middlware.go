package main

import (
	"fmt"
	"net/http"
	"time"
)

// middleware
func loggingMiddleware(next http.Handler) http.Handler {
	//hof
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//Logic Before the handler
		fmt.Printf("Started %s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
		//call "the next hanlder in the chain"
		next.ServeHTTP(w, r)
		//logic after the handler
		w.Header().Add("companyName", "HCL Technologies")
		fmt.Fprintf(w, "Completed in %s", time.Since(start))
	})
}

func main() {
	//create mutex version of server
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//logic
		fmt.Fprintln(w, "Hello,Subramanian")
	})
	//wrap the mux with middlware
	wrappedMux := loggingMiddleware(mux)
	fmt.Println("Server has started....")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

}
