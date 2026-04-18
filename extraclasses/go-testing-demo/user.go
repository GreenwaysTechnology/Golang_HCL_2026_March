package main

// Repository interface allows us to mock the DB
type Repository interface {
    GetUserName(id int) string
}

type UserService struct {
    repo Repository
}

func (s *UserService) WelcomeUser(id int) string {
    name := s.repo.GetUserName(id)
    return "Hello, " + name
}