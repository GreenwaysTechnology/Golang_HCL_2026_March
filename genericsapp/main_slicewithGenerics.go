package main

import "fmt"

type Stack[T any] struct {
	//slices with T
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	n := len(s.items)
	item := s.items[n-1]
	s.items = s.items[:n-1]
	return item
}

func main() {
	stack := Stack[float32]{}
	//var value int64=100
	stack.Push(5)
	stack.Push(3.14)
	//stack.Push(value)
	fmt.Println(stack.Pop())
}
