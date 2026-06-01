package repository

import (
	"banking-system/internal/models"
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err := r.DB.Transaction(func(tx *gorm.DB) error {

		var account models.Account

		result := tx.Where("user_id = ?", userId).First(&account)

		if result.Error != nil {
			return result.Error
		}

		if account.Balance < amount {
			return errors.New("Insufficient Balance")
		}

		result = tx.Model(&models.Account{}).Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Where("id = ?", account.ID).UpdateColumn("balance", gorm.Expr("balance - ?", amount))

		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return "", err
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

func (r *AccountRepository) TransferMoney(userId uuid.UUID, amount int64, fromAccountNumber, toAccountNumber string) (string, error) {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if amount <= 0 {
			return errors.New("invalid amount")
		}

		var from models.Account

		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? AND account_number = ?", userId, fromAccountNumber).
			First(&from).Error; err != nil {
			return err
		}

		if from.Balance < amount {
			return errors.New("insufficient balance")
		}

		if err := tx.Model(&models.Account{}).
			Where("account_number = ?", fromAccountNumber).
			UpdateColumn("balance", gorm.Expr("balance - ?", amount)).
			Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Account{}).
			Where("account_number = ?", toAccountNumber).
			UpdateColumn("balance", gorm.Expr("balance + ?", amount)).
			Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return "Amount transferred successfully", nil
}

func (r *AccountRepository) GetTransactions(userId uuid.UUID) ([]models.Transaction, error) {

	var transactions []models.Transaction

	err := r.DB.Model(&models.Transaction{}).Where("user_id = ?", userId).Find(&transactions)

	if err != nil {
		return []models.Transaction{}, nil
	}

	return transactions, nil
}
