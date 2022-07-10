package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (AccountController) Create(c *gin.Context) {
	account := models.Account{}
	err := c.Bind(&account)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	account.CreateAccount(&account)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   account,
	})
}
