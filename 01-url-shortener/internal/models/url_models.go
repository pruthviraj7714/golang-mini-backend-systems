package models

import "time"

type Url struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ShortURL  string    `json:"short_url" gorm:"not null;unique"`
	LongURL   string    `json:"long_url" gorm:"not null"`
	UserID    int       `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
