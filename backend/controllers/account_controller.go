package controllers

import (
	"encoding/json"
	"net/http"
	"omoshiroGoMoji/backend/models"
)

func AccountCreate(w http.ResponseWriter, r *http.Request) {
	account := models.Account{Email: r.PostFormValue("email"), Password: r.PostFormValue("password")}
	account.Create(&account)
	accountResponse, err := json.Marshal(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(accountResponse)
}
