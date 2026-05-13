package services

import "url-shortener/internal/repository"

type UserServices struct {
	UserRepository *repository.UserRepository
}

func (s *UserServices) RegisterUser(email, password string) error {
	return s.UserRepository.RegisterUser(email, password)
}

func (s *UserServices) LoginUser(email, password string) (string, error) {
	return s.UserRepository.LoginUser(email, password)
}
