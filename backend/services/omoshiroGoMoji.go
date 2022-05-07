package services

import (
	"omoshiroGoMoji/backend/databases"
	"omoshiroGoMoji/backend/models"
)

type OmoshiroGoMojiService struct{}

func (OmoshiroGoMojiService) CreateOmoshiroGoMojiService(omoshiroGoMoji *models.OmoshiroGoMoji) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Create(&omoshiroGoMoji)
	return nil
}

func (OmoshiroGoMojiService) GetOmoshiroGoMojiList() []models.OmoshiroGoMoji {
	var omoshiroGomojis []models.OmoshiroGoMoji
	dbConnection := databases.GetDBConnection()
	dbConnection.Find(&omoshiroGomojis)
	return omoshiroGomojis
}

// func (OmoshiroGoMojiService) UpdateOmoshiroGoMoji(newOmoshiroGoMoji *model.OmoshiroGoMoji) error {
// 	_, err := DbEngine.Id(newOmoshiroGoMoji.Id).Update(newOmoshiroGoMoji)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (OmoshiroGoMojiService) DeleteOmoshiroGoMoji(id int) error {
// 	omoshiroGoMoji := new(model.OmoshiroGoMoji)
// 	_, err := DbEngine.Id(id).Delete(omoshiroGoMoji)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
