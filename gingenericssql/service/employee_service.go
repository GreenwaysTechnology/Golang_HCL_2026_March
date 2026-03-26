package service

import "github.com/gingenericssql/repository"

// Generic Service Object
type Service[T any] struct {
	repo repository.Repository[T]
}

// generic constructor
func NewService[T any](r repository.Repository[T]) *Service[T] {
	return &Service[T]{repo: r}
}

// Generic curd apis
func (s *Service[T]) GetAll() ([]T, error) {
	return s.repo.GetAll()
}
func (s *Service[T]) GetByID(id int) (T, error) {
	return s.repo.GetByID(id)
}

func (s *Service[T]) Create(e T) (T, error) {
	return s.repo.Create(e)
}

func (s *Service[T]) Update(id int, e T) (T, error) {
	return s.repo.Update(id, e)
}

func (s *Service[T]) Delete(id int) error {
	return s.repo.Delete(id)
}
