package repository

import (
	"banking-system/internal/models"
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func generateRandomAccountNumber() string {
	return fmt.Sprintf("ACC-%d", rand.Intn(10000000000))
}

func (r *AccountRepository) CreateAccount(userId uuid.UUID) (uuid.UUID, error) {

	account := models.Account{
		AccountNumber: generateRandomAccountNumber(),
		UserID:        userId,
		Balance:       1000,
		Status:        models.AccountStatusActive,
	}

	response := r.DB.Create(&account)

	if response.Error != nil {
		return uuid.UUID{}, response.Error
	}

	return account.ID, nil
}

func (r *AccountRepository) GetAccount(userId uuid.UUID) (*models.Account, error) {
	var account models.Account

	err := r.DB.Where("user_id = ?", userId).First(&account).Error

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) WithdrawMoney(userId uuid.UUID, amount int64) (string, error) {
	var account models.Account

	err := r.DB.Where("user_id = ?", userId).First(&account).Error

	if err != nil {
		return "", err
	}

	fmt.Print(account)

	if account.Balance < amount {
		return "", errors.New("Insufficient Balance")
	}

	result := r.DB.Model(&models.Account{}).Where("id = ?", account.ID).UpdateColumn("balance", gorm.Expr("balance - ?", amount))

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", errors.New("no rows updated")
	}

	return "Amount successfully Withdrawn", nil
}

func (r *AccountRepository) DepositMoney(userId uuid.UUID, amount int64) (string, error) {
	var account models.Account

	err := r.DB.Model(&models.Account{}).Where("user_id = ?", userId).First(&account).Error

	if err != nil {
		return "", err
	}

	result := r.DB.Model(&models.Account{}).Where("id = ?", account.ID).UpdateColumn("balance", gorm.Expr("balance + ?", amount))

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", errors.New("no rows updated")
	}

	return "Amount successfully Deposited", nil
}
