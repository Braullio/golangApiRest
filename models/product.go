package models

import (
	"time"
)

type Product struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
