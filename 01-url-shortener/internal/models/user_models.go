package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"auto_create_time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"auto_update_time"`
}
