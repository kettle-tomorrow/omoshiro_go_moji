package services

import (
	"omoshiroGoMoji/backend/databases"
	"omoshiroGoMoji/backend/models"

	"golang.org/x/crypto/bcrypt"
)

type AccountService struct{}

func (AccountService) GetAccountByEmail(email string) models.Account {
	var account models.Account
	dbConnection := databases.GetDBConnection()
	dbConnection.Where("email = ?", email).First(&account)
	return account
}

func (AccountService) CreateAccount(account *models.Account) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)
	dbConnection := databases.GetDBConnection()
	dbConnection.Create(account)
	return nil
}
