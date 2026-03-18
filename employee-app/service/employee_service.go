package service

import (
	"github.com/employee/model"
	"github.com/employee/repository"
)

type EmployeeService struct {
	//Depedency Injection - Composition/Embeding
	repo repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: r}
}

// apis
func (s *EmployeeService) Create(emp model.Employee) error {
	return s.repo.Create(emp)
}
func (s *EmployeeService) GetAll() ([]model.Employee, error) {
	return s.repo.GetAll()
}
func (s *EmployeeService) GetById(id int) (*model.Employee, error) {
	return s.repo.GetByID(id)
}
func (s *EmployeeService) Update(emp model.Employee) error {
	return s.repo.Update(emp)
}
func (s *EmployeeService) Delete(id int) error {
	return s.repo.Delete(id)
}
