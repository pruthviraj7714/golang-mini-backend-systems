package models

import "github.com/google/uuid"

type TransactionType string

const (
	Transfer TransactionType = "TRANSFER"
	Withdraw TransactionType = "WITHDRAW"
	Deposit  TransactionType = "DEPOSIT"
)

type Transaction struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	SenderAccountNumber string `json:"sender_account_number" gorm:unique;index"`

	ReceiverAccountNumber string `json:"receiver_account_number" gorm:unique;index"`

	Type TransactionType `json:"type"`
}
