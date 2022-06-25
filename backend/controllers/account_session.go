package controllers

import (
	"encoding/json"
	"net/http"
	"omoshiroGoMoji/backend/models"
	"omoshiroGoMoji/backend/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	requestAccount := models.Account{}
	err := c.Bind(&requestAccount)
	if err != nil {
		c.String(http.StatusBadRequest, "login failed1")
	} else {
		accountService := services.AccountService{}
		account := accountService.GetAccountByEmail(requestAccount.Email)

		err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(requestAccount.Password))
		if err != nil {
			c.String(http.StatusBadRequest, "login failed2")
		} else {
			session := sessions.Default(c)
			loginUser, err := json.Marshal(account)
			if err == nil {
				session.Set("loginUser", string(loginUser))
				session.Save()
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusInternalServerError)
			}
		}
	}
}
