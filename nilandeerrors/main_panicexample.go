package main

import "fmt"

func throwPanic() {
	fmt.Println("Start")
	panic("Something bad happened")
	fmt.Println("something is going on")
	fmt.Println("end")
}

// nil pointer panic
/**
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0x7ff78bd149b6]

goroutine 1 [running]:
//main.NilPointerPanic()
        G:/HCL/2026/march/nilandeerrors/main_panicexample.go:15 +0x16
main.main()
        G:/HCL/2026/march/nilandeerrors/main_panicexample.go:20 +0xf
exit status 2
*/
func NilPointerPanic() {
	var p *int
	//var x = 10
	//p = &x
	//now this variable is not pointing any address you get panic(Exception in Java)
	fmt.Println(*p)
}

func main() {
	//throwPanic()
	NilPointerPanic()
	fmt.Println("continue")
}
