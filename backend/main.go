package main

import (
	"encoding/json"
	"net/http"
	"omoshiroGoMoji/backend/controllers"
	"omoshiroGoMoji/backend/databases"
	"omoshiroGoMoji/backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func main() {
	databases.Init(&models.User{}, &models.OmoshiroGoMoji{}, &models.Account{})
	defer databases.Close()
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	// セッションCookieの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/login", controllers.Login)
	apiV1.POST("/account", controllers.AccountCreate)

	omoshiroGoMojiRouter := apiV1.Group("/omoshiro_go_moji")
	omoshiroGoMojiRouter.GET("/list", controllers.OmoshiroGoMojiList)
	omoshiroGoMojiRouter.GET("/:id", controllers.OmoshiroGoMojiShow)
	omoshiroGoMojiRouter.Use(loginCheckMiddleware())
	omoshiroGoMojiRouter.POST("", controllers.OmoshiroGoMojiCreate)
	omoshiroGoMojiRouter.PATCH("/:id", controllers.OmoshiroGoMojiUpdate)
	omoshiroGoMojiRouter.DELETE("/:id", controllers.OmoshiroGoMojiDelete)

	user := apiV1.Group("/user")
	user.Use(loginCheckMiddleware())
	user.GET("/list", controllers.UserList)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to おもしろGo文字",
		})
	})
	router.Run(":3000")
}

func loginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// Json文字列がinterdace型で格納されておりdproxyのライブラリを使用して値を取り出す
		loginUserJson, getLoginUserErr := dproxy.New(session.Get("loginUser")).String()

		if getLoginUserErr != nil {
			c.String(http.StatusUnauthorized, "unauthorized1")
			c.Abort()
		}

		var loginInfo models.Account
		jsonUnmarshalErr := json.Unmarshal([]byte(loginUserJson), &loginInfo)
		if jsonUnmarshalErr != nil {
			c.String(http.StatusUnauthorized, "unauthorized2")
			c.Abort()
		}

		c.Set("currentUserID", loginInfo.ID)
		c.Next()
	}
}
