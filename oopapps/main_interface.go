package main

import "fmt"

// interface declaration
type Speaker interface {
	//method declaration
	Speak()
}

// any type that has Speak mehthod satisifes the Speaker interface
type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Dog barks")
}
func (d Dog) Eat() {
	fmt.Println("Dog eats")
}

type Cat struct {
}

func (c Cat) Speak() {
	fmt.Println("Cat memows")
}

func main() {
	//variable with interface declaration
	var s Speaker
	//now s type is Speaker and Dog is implementor
	s = Dog{}
	s.Speak()
	//this method sepcific to struct not common method
	//Type assertion : is used to  extract the concrete value from an interface
	//value:=interfaceVariable.(*ConcreteType | concreteType)
	//
	var tmp = s.(Dog)
	tmp.Eat()
	s = Cat{}
	s.Speak()
}
