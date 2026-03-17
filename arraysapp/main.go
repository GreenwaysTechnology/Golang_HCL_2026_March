package main

import "fmt"

type Data struct {
	Values [3]int
}
type User struct {
	Name string
}

func main() {
	d := Data{
		Values: [3]int{1, 2, 3},
	}
	fmt.Println(d.Values)
	users := [2]User{
		{"John"},
		{"Paul"},
	}
	for _, user := range users {
		fmt.Println(user.Name)
	}
}
