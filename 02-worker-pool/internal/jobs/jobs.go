package jobs

import "time"

type Status string

const (
	pending    Status = "pending"
	processing Status = "processing"
	completed  Status = "completed"
	failed     Status = "failed"
)

type Job struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Type        string    `json:"type" gorm:"not null"`
	Payload     any       `json:"payload" gorm:"type:json"`
	Status      Status    `json:"status" gorm:"default:pending"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CompletedAt time.Time `json:"completed_at" gorm:"null"`
}
