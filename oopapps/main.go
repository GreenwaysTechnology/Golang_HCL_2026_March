package main

import "fmt"

type Customer struct {
	FirstName string
	LastName  string
}

// value receiver: never updates instance members
func (c Customer) updateName() {
	c.FirstName = "Subramanian"
	c.LastName = "Murugan"
}

// pointer/refence receiver updates instance members, because pointers share memory
func (c *Customer) updateNameUsingReference() {
	c.FirstName = "Subramanian"
	c.LastName = "Murugan"
}
func main() {
	customer := Customer{}
	fmt.Println("Before updateName")
	fmt.Println(customer.FirstName, customer.LastName)
	//customer.updateName()
	customer.updateNameUsingReference()
	fmt.Println("After updateName")
	fmt.Println(customer.FirstName, customer.LastName)
}
