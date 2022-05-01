package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db_connection := DBConnect()
	db, _ := db_connection.DB()
	defer db.Close()
	DBMigrate(db_connection)
	engine := gin.Default()
	// apiV1 := engine.Group("/api/v1")
	// {
	// 	omoshiroGoMojiEngine := apiV1.Group("/omoshiro-go-moji")
	// 	{
	// 		omoshiroGoMojiEngine.GET("/list", controller.OmoshiroGoMojiList)
	// 		omoshiroGoMojiEngine.POST("/add", controller.OmoshiroGoMojiAdd)
	// 		omoshiroGoMojiEngine.PUT("/update", controller.OmoshiroGoMojiUpdate)
	// 		omoshiroGoMojiEngine.DELETE("/delete", controller.OmoshiroGoMojiDelete)
	// 	}
	// }
	engine.GET("/", func(c *gin.Context) {
		var user User
		db_connection.Find(&user)
		fmt.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to おもしろGo文字",
			"data":    user,
		})
	})
	engine.Run(":3000")
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"Name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}

func DBConnect() *gorm.DB {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "omoshiro_go_moji_development"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
