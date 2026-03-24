package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetPosts() {
	url := "https://jsonplaceholder.typicode.com/posts"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
func main() {
	GetPosts()
}
