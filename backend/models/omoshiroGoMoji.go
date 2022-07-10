package models

import (
	"omoshiroGoMoji/backend/databases"
	"time"
)

type OmoshiroGoMoji struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name" form:"name" binding:"required,len=5"`
	UserID    uint
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (OmoshiroGoMoji) GetOmoshiroGoMojiList() []OmoshiroGoMoji {
	var omoshiroGomojis []OmoshiroGoMoji
	dbConnection := databases.GetDBConnection()
	dbConnection.Find(&omoshiroGomojis)
	return omoshiroGomojis
}

func (OmoshiroGoMoji) GetOmoshiroGoMoji(id string) OmoshiroGoMoji {
	var omoshiroGoMoji OmoshiroGoMoji
	dbConnection := databases.GetDBConnection()
	dbConnection.First(&omoshiroGoMoji, id)
	return omoshiroGoMoji
}

func (OmoshiroGoMoji) CreateOmoshiroGoMoji(omoshiroGoMoji *OmoshiroGoMoji) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Create(omoshiroGoMoji)
	return nil
}

func (OmoshiroGoMoji) UpdateOmoshiroGoMoji(omoshiroGoMoji *OmoshiroGoMoji) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Save(omoshiroGoMoji)
	return nil
}

func (OmoshiroGoMoji) DeleteOmoshiroGoMoji(id string) error {
	dbConnection := databases.GetDBConnection()
	dbConnection.Delete(OmoshiroGoMoji{}, id)
	return nil
}
