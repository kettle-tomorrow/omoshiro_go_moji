package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Index(c *gin.Context) {
	user := models.User{}
	userList := user.List()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    userList,
	})
}
