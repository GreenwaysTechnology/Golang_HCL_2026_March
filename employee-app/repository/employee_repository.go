package repository

import "github.com/employee/model"

type EmployeeRepository interface {
	Create(emp model.Employee) error
	GetAll() ([]model.Employee, error)
	GetByID(id int) (*model.Employee, error)
	Update(emp model.Employee) error
	Delete(id int) error
}
