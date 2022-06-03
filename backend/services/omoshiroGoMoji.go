package services

import (
	"omoshiroGoMoji/backend/databases"
	"omoshiroGoMoji/backend/models"
)

type OmoshiroGoMojiService struct{}

func (OmoshiroGoMojiService) CreateOmoshiroGoMojiService(omoshiroGoMoji *models.OmoshiroGoMoji) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Create(omoshiroGoMoji)
	return nil
}

func (OmoshiroGoMojiService) GetOmoshiroGoMojiList() []models.OmoshiroGoMoji {
	var omoshiroGomojis []models.OmoshiroGoMoji
	dbConnection := databases.GetDBConnection()
	dbConnection.Find(&omoshiroGomojis)
	return omoshiroGomojis
}

func (OmoshiroGoMojiService) GetOmoshiroGoMoji(id string) models.OmoshiroGoMoji {
	var omoshiroGoMoji models.OmoshiroGoMoji
	dbConnection := databases.GetDBConnection()
	dbConnection.First(&omoshiroGoMoji, id)
	return omoshiroGoMoji
}

func (OmoshiroGoMojiService) UpdateOmoshiroGoMoji(omoshiroGoMoji *models.OmoshiroGoMoji) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Save(omoshiroGoMoji)
	return nil
}

func (OmoshiroGoMojiService) DeleteOmoshiroGoMoji(id string) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Delete(models.OmoshiroGoMoji{}, id)
	return nil
}
