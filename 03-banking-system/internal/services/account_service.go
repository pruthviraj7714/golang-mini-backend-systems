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

func (s *AccountService) WithdrawMoney(userId uuid.UUID, amount int64) (string, error) {
	return s.Repo.WithdrawMoney(userId, amount)
}

func (s *AccountService) DepositMoney(userId uuid.UUID, amount int64) (string, error) {
	return s.Repo.DepositMoney(userId, amount)
}
