package services

import "banking-system/internal/repository"

type UserService struct {
	Repo       *repository.UserRepository
	JWTService *JWTService
}

func (s *UserService) RegisterUser(email, password string) error {
	return s.Repo.Register(email, password)
}

func (s *UserService) LoginUser(email, password string) (string, string, error) {
	return s.Repo.Login(email, password, s.JWTService.AccessSecret, s.JWTService.RefreshSecret)
}
