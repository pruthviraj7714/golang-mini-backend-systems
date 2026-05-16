package services

import "url-shortener/internal/repository"

type UserService struct {
	UserRepository *repository.UserRepository
}

func (s *UserService) RegisterUser(email, password string) error {
	return s.UserRepository.RegisterUser(email, password)
}

func (s *UserService) LoginUser(email, password string) (string, error) {
	return s.UserRepository.LoginUser(email, password)
}
