package model

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	ID         int            `json:"id"`
	ClientID   uint           `json:"client_id"`
	Address    string         `json:"address"`
	PostalCode string         `json:"postal_code"`
	City       string         `json:"city"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
