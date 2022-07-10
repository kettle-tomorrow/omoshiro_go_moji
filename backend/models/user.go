package models

import (
	"omoshiroGoMoji/backend/databases"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name" form:"name" binding:"required,max=10"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (User) List() []User {
	var users []User
	dbConnection := databases.GetDBConnection()
	dbConnection.Find(&users)
	return users
}
