package main

import (
	"fmt"
	"net/http"
	"time"
)

// middleware
func LogginMiddleware(next http.Handler) http.Handler {
	//hof
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//Logic Before the handler
		fmt.Printf("Started %s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
		//call "the next hanlder in the chain"
		w.Header().Add("companyName", "HCL Technologies")
		next.ServeHTTP(w, r)
		//logic after the handler
		fmt.Fprintf(w, "Completed in %s", time.Since(start))
	})
}

// AuthMiddleware Auth middlware : checks for a specific header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")
		if token != "secret-pass" {
			fmt.Println("[Auth] Unauthorized Attempt")
			http.Error(w, "Forbidden: Invalid Token", http.StatusForbidden)
			return // stop the chain here
		}
		fmt.Println("[Auth] User Authenticated")
		next.ServeHTTP(w, r)

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
	//order-->Request--->Logging--->Auth--Mux--Handler
	//middleware chain a(b(c(d(e(mux)))))
	wrappedHandler := LogginMiddleware(AuthMiddleware(mux))
	fmt.Println("Server has started....")
	err := http.ListenAndServe(":8080", wrappedHandler)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

}
