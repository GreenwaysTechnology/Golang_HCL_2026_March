package main

import (
	"github.com/employee/controller"
	"github.com/employee/repository"
	"github.com/employee/service"
)

func main() {
	repo := repository.NewFileRepository("storage/employees.json")
	service := service.NewEmployeeService(repo)
	handler := controller.NewCLIHandler(service)
	handler.Start()
}
