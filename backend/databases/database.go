package databases

import (
	"omoshiroGoMoji/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func Init() {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "omoshiro_go_moji_development"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	var err error
	dbConnection, err = gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	dbConnection.AutoMigrate(&models.User{})
}

func GetDBConnection() *gorm.DB {
	return dbConnection
}

func Close() {
	db, _ := dbConnection.DB()
	db.Close()
}
