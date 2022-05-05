package services

import (
	"omoshiro_go_moji/backend/databases"
	"omoshiro_go_moji/backend/models"
)

type UserService struct{}

func (UserService) GetUserList() []models.User {
	var users []models.User
	db_connection := databases.GetDBConnection()
	db_connection.Find(&users)
	return users
}
