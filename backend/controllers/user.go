package controllers

import (
	"net/http"
	service "omoshiroGoMoji/backend/services"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	userService := service.UserService{}
	userList := userService.GetUserList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    userList,
	})
}
