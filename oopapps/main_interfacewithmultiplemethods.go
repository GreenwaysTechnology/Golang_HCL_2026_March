package main

import "fmt"

// interface with multiple methods
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return c.Radius * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * c.Radius
}
func printShape(shape Shape) {
	fmt.Println("Area", shape.Area())
	fmt.Println("Perimeter", shape.Perimeter())
}

func main() {
	//passing instance of Rectangle to the Type called Shape
	printShape(Rectangle{Width: 10, Height: 5})
	printShape(Circle{Radius: 5})
}
