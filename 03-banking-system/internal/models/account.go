package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "ACTIVE"
	AccountStatusSuspended AccountStatus = "SUSPENDED"
	AccountStatusClosed    AccountStatus = "CLOSED"
)

type Account struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	// Account Number is the unique identifier for the account - used for deposits, withdrawals, transfers
	AccountNumber string `json:"account_number" gorm:"not null;unique;index"`

	// UserID is the foreign key linking the account to a user
	UserID uuid.UUID `json:"user_id" gorm:"not null"`

	// Balance is the current balance in the account
	Balance int64 `json:"balance" gorm:"not null;default:0"`

	// Status is the status of the account (e.g., "ACTIVE", "SUSPENDED", "CLOSED")
	Status AccountStatus `json:"status" gorm:"not null;default:'ACTIVE'"`

	// CreatedAt is the timestamp when the account was created
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// UpdatedAt is the timestamp when the account was last updated
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// User represents the owner of the account
	User User `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
