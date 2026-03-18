package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/employee/model"
)

type FileRepository struct {
	filePath string
}

// Faactory function
func NewFileRepository(filePath string) *FileRepository {
	return &FileRepository{filePath: filePath}
}

// load json file
func (f *FileRepository) load() ([]model.Employee, error) {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Employee{}, nil
		}
		return nil, err
	}
	//slice
	var employees []model.Employee
	if len(file) == 0 {
		return []model.Employee{}, nil
	}
	err = json.Unmarshal(file, &employees)
	return employees, err
}

// save data into json file
func (f *FileRepository) save(data []model.Employee) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, bytes, 0644)
}

// Create
func (f *FileRepository) Create(emp model.Employee) error {
	employees, _ := f.load()
	for _, e := range employees {
		if e.ID == emp.ID {
			return fmt.Errorf("employee with id %d already exists", emp.ID)
		}
	}
	employees = append(employees, emp)
	return f.save(employees)
}
func (f *FileRepository) GetByID(id int) (*model.Employee, error) {
	employees, _ := f.load()
	for _, e := range employees {
		if e.ID == id {
			return &e, nil
		}
	}
	//return nil,errors.New("Employee not found")
	return nil, fmt.Errorf("employee with id %d does not exist", id)
}

// Update
func (f *FileRepository) Update(emp model.Employee) error {
	employees, _ := f.load()
	for i, e := range employees {
		if e.ID == emp.ID {
			employees[i] = emp
			return f.save(employees)
		}

	}
	return errors.New("employee with id does not exist")
}
func (f *FileRepository) GetAll() ([]model.Employee, error) {
	return f.load()
}
func (f *FileRepository) Delete(id int) error {
	employees, _ := f.load()
	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return f.save(employees)
		}
	}
	return errors.New("employee with id does not exist")
}
