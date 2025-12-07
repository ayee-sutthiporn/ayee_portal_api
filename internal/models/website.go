package models

import (
	"time"

	"github.com/google/uuid"
)

type Website struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	URL         string    `gorm:"not null" json:"url"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `gorm:"type:uuid" json:"categoryId"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Icon        string    `json:"icon"`
	IsVisible   bool      `gorm:"default:true" json:"isVisible"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
