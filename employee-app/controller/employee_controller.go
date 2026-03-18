package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/employee/model"
	"github.com/employee/service"
)

type CLIHandler struct {
	//dependency
	service *service.EmployeeService
	reader  *bufio.Reader
}

// create Controller instance
func NewCLIHandler(s *service.EmployeeService) *CLIHandler {
	return &CLIHandler{
		service: s,
		reader:  bufio.NewReader(os.Stdin), //you bind the keyword with buffer
	}
}

// methods of controller
func (h *CLIHandler) ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := h.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// read int
func (h *CLIHandler) ReadInt(prompt string) int {
	for {
		val := h.ReadInput(prompt)
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid Number")
			continue
		}
		return num
	}
}
func (h *CLIHandler) create() {
	id := h.ReadInt("Enter Employee ID :")
	name := h.ReadInput("Enter Employee Name :")
	age := h.ReadInt("Enter Employee Age :")
	err := h.service.Create(model.Employee{ID: id, Name: name, Age: age})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created Employee successfully")
}
func (h *CLIHandler) list() {
	data, _ := h.service.GetAll()
	for _, d := range data {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", d.ID, d.Name, d.Age)
	}
}
func (h *CLIHandler) get() {
	id := h.ReadInt("Enter Employee ID :")
	e, err := h.service.GetById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", e.ID, e.Name, e.Age)
}
func (h *CLIHandler) update() {
	id := h.ReadInt("Enter Employee ID :")
	name := h.ReadInput("Enter Employee Name :")
	age := h.ReadInt("Enter Employee Age :")
	err := h.service.Update(model.Employee{ID: id, Name: name, Age: age})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d Employee successfully", id)
}
func (h *CLIHandler) delete() {
	id := h.ReadInt("Enter Employee ID :")
	err := h.service.Delete(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d Employee successfully", id)
}

func (h *CLIHandler) Start() {
	for {
		fmt.Println("\n1.Add 2.List 3.Get 4.Update 5.Delete 6.Exit")
		choice := h.ReadInt("Enter your choice: ")
		switch choice {
		case 1:
			h.create()
		case 2:
			h.list()
		case 3:
			h.get()
		case 4:
			h.update()
		case 5:
			h.delete()
		case 6:
			fmt.Println("Thank You!")
			os.Exit(0)
		}
	}
}
