package repository

import (
	"gormginapp/model"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}
func (r *EmployeeRepository) Create(emp *model.Employee) error {
	return r.db.Create(emp).Error
}
func (r *EmployeeRepository) GetByID(id uint) (*model.Employee, error) {
	var emp model.Employee
	err := r.db.First(&emp, id).Error
	return &emp, err
}
func (r *EmployeeRepository) GetAll(limit, offset int) ([]model.Employee, error) {
	var emps []model.Employee
	err := r.db.Limit(limit).Offset(offset).Find(&emps).Error
	return emps, err
}
func (r *EmployeeRepository) Update(emp *model.Employee) error {
	return r.db.Save(emp).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	return r.db.Delete(&model.Employee{}, id).Error
}
