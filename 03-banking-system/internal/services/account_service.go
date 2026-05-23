package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repository"

	"github.com/google/uuid"
)

type AccountService struct {
	Repo *repository.AccountRepository
}

func (s *AccountService) CreateAccount(userId uuid.UUID) (uuid.UUID, error) {
	return s.Repo.CreateAccount(userId)
}

func (s *AccountService) GetAccount(userId uuid.UUID) (*models.Account, error) {
	return s.Repo.GetAccount(userId)
}
