package models

import (
	"time"
)

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null" json:"email" form:"email" binding:"required,max=255"`
	Password  string    `gorm:"not null" json:"password" form:"password" binding:"required,max=100"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
