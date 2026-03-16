package main

import "fmt"

// interface embding, go does not support inheritance, composition
type Reader interface {
	Read()
}
type Writer interface {
	Write()
}
type ReadWriter interface {
	Reader
	Writer
}
type File struct {
	Name string
}

func (f File) Read() {
	fmt.Println("Read", f.Name)
}
func (f File) Write() {
	fmt.Println("Write", f.Name)
}
func Process(rw ReadWriter) {
	rw.Read()
	rw.Write()
}
func main() {
	//file := File{Name: "data.txt"}
	//var rw ReadWriter = file
	//var rw ReadWriter = File{Name: "test"}
	var rw ReadWriter = File{Name: "test"}
	rw.Read()
	rw.Write()
	Process(rw)
}
