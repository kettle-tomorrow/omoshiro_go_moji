package databases

import (
	"omoshiro_go_moji/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_connection *gorm.DB

func Init() {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "omoshiro_go_moji_development"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	var err error
	db_connection, err = gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db_connection.AutoMigrate(&models.User{})
}

func GetDBConnection() *gorm.DB {
	return db_connection
}

func Close() {
	db, _ := db_connection.DB()
	db.Close()
}
