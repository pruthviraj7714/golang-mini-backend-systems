package models

import (
	"time"
)

type NotificationStatus string

var (
	Pending NotificationStatus = "pending"
	Sent    NotificationStatus = "sent"
	Failed  NotificationStatus = "failed"
)

type Notification struct {
	ID        int64              `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string             `json:"user_id" gorm:"index"`
	Message   string             `json:"message"`
	Status    NotificationStatus `json:"status" gorm:"default:'pending'"`
	CreatedAt time.Time          `json:"created_at" gorm:"autoCreateTime"`
}
