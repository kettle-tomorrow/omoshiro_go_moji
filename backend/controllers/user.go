package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := models.User{}
	userList := user.GetUserList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    userList,
	})
}
