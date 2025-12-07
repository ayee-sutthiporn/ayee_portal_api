package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name  string    `gorm:"not null" json:"name"`
	Icon  string    `json:"icon"`
	Order int       `gorm:"default:0" json:"order"`
}
