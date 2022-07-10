package models

import (
	"omoshiroGoMoji/backend/databases"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null" json:"email" form:"email" binding:"required,max=255"`
	Password  string    `gorm:"not null" json:"password" form:"password" binding:"required,max=100"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (Account) GetByEmail(email string) Account {
	var account Account
	dbConnection := databases.GetDBConnection()
	dbConnection.Where("email = ?", email).First(&account)
	return account
}

func (Account) Create(account *Account) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)
	dbConnection := databases.GetDBConnection()
	dbConnection.Create(account)
	return nil
}
