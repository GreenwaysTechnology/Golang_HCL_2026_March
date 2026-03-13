package main

import "fmt"

// class declaration - structure
type Person struct {
	Name  string
	Age   int
	ID    string
	EMAIL string
}

func main() {
	//Basic instance creation
	var p Person
	p.Name = "Subramanian Murugan"
	p.Age = 46
	p.ID = "123456"
	p.EMAIL = "sasubramanian_md@hotmail.com"

	fmt.Println(p)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.ID)
	fmt.Println(p.EMAIL)

	//Initalize using literal pattern
	var p2 Person = Person{
		Name:  "Subramanian Murugan",
		Age:   42,
		ID:    "123456",
		EMAIL: "sasubramanian_md@hotmail.com",
	}
	fmt.Println(p2)

	//without field names
	var p3 Person = Person{"Subramanian Murugan",
		42, "123456",
		"sasubramanian_md@gmail.com"}
	fmt.Println(p3)

}
