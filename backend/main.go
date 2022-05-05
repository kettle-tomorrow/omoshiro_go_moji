package main

import (
	"fmt"
	"net/http"
	"omoshiro_go_moji/backend/databases"
	"omoshiro_go_moji/backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.Init()
	defer databases.Close()
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
		var user models.User
		db_connection := databases.GetDBConnection()
		db_connection.Find(&user)
		fmt.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to おもしろGo文字",
			"data":    user,
		})
	})
	engine.Run(":3000")
}
