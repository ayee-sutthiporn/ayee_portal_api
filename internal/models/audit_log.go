package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    string    `json:"userId"`
	Username  string    `json:"username"`
	Action    string    `json:"action"` // e.g., 'CREATE_WEBSITE'
	Details   string    `json:"details"`
	Timestamp time.Time `json:"timestamp"`
}
