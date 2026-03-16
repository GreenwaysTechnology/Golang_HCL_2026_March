package main

var empty interface{}

func main() {
	empty = 10
	//it prints some address because it is reference type
	v, ok := empty.(int)
	if ok {
		println(v)
	}
	//println(v, ok)
	//empty = "Hello World"
	//println(empty)
	//empty = 23.80
	//println(empty)
}
