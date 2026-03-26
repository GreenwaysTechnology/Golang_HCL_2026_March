package main

import "fmt"

type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

type User struct {
	Name string
}

func (u User) String() string {
	return "Name:" + u.Name
}

type MyInt int

func (u MyInt) String() string {
	return fmt.Sprintf("%d", u)
}

func main() {
	num := MyInt(1)
	PrintString(num)
	person := User{"John"}
	PrintString(person)
}
