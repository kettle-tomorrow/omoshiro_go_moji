package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"Name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
