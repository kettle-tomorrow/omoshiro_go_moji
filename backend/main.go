package main

import (
	"net/http"
	"omoshiroGoMoji/backend/controllers"
	"omoshiroGoMoji/backend/databases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	databases.Init()
	defer databases.Close()
	engine := gin.Default()

	// CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	apiV1 := engine.Group("/api/v1")
	omoshiroGoMojiEngine := apiV1.Group("/omoshiro_go_moji")
	omoshiroGoMojiEngine.GET("/list", controllers.OmoshiroGoMojiList)
	omoshiroGoMojiEngine.GET("/:id", controllers.OmoshiroGoMojiShow)
	omoshiroGoMojiEngine.POST("", controllers.OmoshiroGoMojiCreate)
	omoshiroGoMojiEngine.PATCH("/:id", controllers.OmoshiroGoMojiUpdate)
	omoshiroGoMojiEngine.DELETE("/:id", controllers.OmoshiroGoMojiDelete)
	user := apiV1.Group("/user")
	user.GET("/list", controllers.UserList)
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to おもしろGo文字",
		})
	})
	engine.Run(":3000")
}
