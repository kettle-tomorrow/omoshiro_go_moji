package main

import (
	"net/http"
	"omoshiroGoMoji/backend/controllers"
	"omoshiroGoMoji/backend/databases"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.Init()
	defer databases.Close()
	engine := gin.Default()
	apiV1 := engine.Group("/api/v1")
	omoshiroGoMojiEngine := apiV1.Group("/omoshiro_go_moji")
	omoshiroGoMojiEngine.GET("/list", controllers.OmoshiroGoMojiList)
	omoshiroGoMojiEngine.POST("/", controllers.OmoshiroGoMojiCreate)
	// 	omoshiroGoMojiEngine.PUT("/update", controller.OmoshiroGoMojiUpdate)
	// 	omoshiroGoMojiEngine.DELETE("/delete", controller.OmoshiroGoMojiDelete)
	user := apiV1.Group("/user")
	user.GET("/list", controllers.UserList)
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to おもしろGo文字",
		})
	})
	engine.Run(":3000")
}
