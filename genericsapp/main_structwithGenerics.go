package main

import "fmt"

type Box[T any] struct {
	Value T
}

func (b Box[T]) Get() T {
	return b.Value
}
func (b *Box[T]) Set(newValue T) {
	b.Value = newValue
}
func main() {
	//Using box with Integer
	intBox := Box[int]{1000}
	intBox.Set(100)
	fmt.Println(intBox.Get())
	floatBox := Box[float32]{1.222}
	fmt.Println(floatBox.Get())
}
