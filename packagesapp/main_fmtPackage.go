package main

import (
	"fmt"
	"os"
)

//fmt package
/**
	fmt is formatted i/o
	1.it prints outputs to console
	2.format strings
	3.read user input
    4.build formated strings
*/

func fmtPrint() {
	//without new line
	fmt.Print("Hello World")
	fmt.Print("how are you?")
}
func fmtPrintLn() {
	//without new line
	fmt.Println("Hello World")
	fmt.Println("how are you?")
}
func fmtPrintF() {
	//if you want to format control using verbs
	name := "Subramanian"
	age := 46
	fmt.Printf("%s is %d years old\n", name, age)
	/**
	%d integer
	%s string
	%t boolean
	%F float\
	%v default value format
	*/
}

// string formating functions
func stringFormating() string {
	name := "subramanian"
	age := 46
	//str:=fmt.Sprintf("Hello %s",name)
	//return str
	return fmt.Sprintf("%s is %d years old\n", name, age)
}

// string concation function
func stringConcat() string {
	//return fmt.Sprint("Hello", "Subrmanian Murugan", 46)
	return fmt.Sprintln("Hello", "Subrmanian Murugan", 46)
}

// getting input from the keyboard
func readInput() (string, int) {
	var input string
	var age int
	fmt.Println("What is your name?")
	//	fmt.Scan(&input)
	//fmt.Scanln(&input)
	fmt.Scanf("%s %d", &input, &age)
	//return input
	return input, age

}

//writes to other outputs (files,network sockets,buffers etc...,console)

func writeToMany() {
	name := "Subramanian"
	fmt.Fprintf(os.Stdout, "Hello %s", name)
}

func main() {
	//fmtPrint()
	//fmtPrintLn()
	//fmtPrintF()
	//fmt.Println(stringFormating())
	//fmt.Println(stringConcat())
	//fmt.Printf("Your Name is %s %d", readInput())
	//name, age := readInput()
	//fmt.Printf("%s is %d years old\n", name, age)
}
