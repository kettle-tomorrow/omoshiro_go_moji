package controllers

import (
	"net/http"
	service "omoshiro_go_moji/backend/services"

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
