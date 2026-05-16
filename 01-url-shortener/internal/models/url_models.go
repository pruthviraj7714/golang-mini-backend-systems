package models

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ShortURL  string    `json:"short_url" gorm:"not null;unique"`
	LongURL   string    `json:"long_url" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null;type:uuid;column:user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
