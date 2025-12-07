package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;type:varchar(36)" json:"id"` // Keycloak UUID
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Role      string    `gorm:"default:'user'" json:"role"` // 'admin' | 'user'
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
