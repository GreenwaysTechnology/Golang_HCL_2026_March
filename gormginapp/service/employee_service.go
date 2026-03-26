package service

import (
	"gormginapp/model"
	"gormginapp/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(r *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: r}
}
func (s *EmployeeService) Create(emp *model.Employee) error {
	return s.repo.Create(emp)
}

func (s *EmployeeService) GetByID(id uint) (*model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) GetAll(limit, offset int) ([]model.Employee, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *EmployeeService) Update(emp *model.Employee) error {
	return s.repo.Update(emp)
}

func (s *EmployeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
