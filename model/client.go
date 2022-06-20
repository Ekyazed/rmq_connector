package model

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        int            `json:"id"`
	UUID      string         `json:"uuid"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
