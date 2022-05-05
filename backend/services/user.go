package services

import (
	"omoshiroGoMoji/backend/databases"
	"omoshiroGoMoji/backend/models"
)

type UserService struct{}

func (UserService) GetUserList() []models.User {
	var users []models.User
	dbConnection := databases.GetDBConnection()
	dbConnection.Find(&users)
	return users
}
