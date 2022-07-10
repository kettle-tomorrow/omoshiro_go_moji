package controllers

import (
	"encoding/json"
	"net/http"
	"omoshiroGoMoji/backend/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AccountSessionController struct{}

func (AccountSessionController) Create(c *gin.Context) {
	requestAccount := models.Account{}
	bindErr := c.Bind(&requestAccount)
	if bindErr != nil {
		c.String(http.StatusBadRequest, "login failed1")
		return
	}

	account := models.Account{}
	account = account.GetAccountByEmail(requestAccount.Email)
	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(requestAccount.Password))
	if comparePasswordErr != nil {
		c.String(http.StatusBadRequest, "login failed2")
		return
	}

	session := sessions.Default(c)
	loginUser, jsonMarshalErr := json.Marshal(account)
	if jsonMarshalErr == nil {
		session.Set("loginUser", string(loginUser))
		session.Save()
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusInternalServerError)
}
