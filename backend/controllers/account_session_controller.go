package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func AccountSessionCreate(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	account = account.GetByEmail(r.PostFormValue("email"))

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(r.PostFormValue("password")))
	if comparePasswordErr != nil {
		http.Error(w, comparePasswordErr.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    strconv.FormatUint(uint64(account.ID), 10),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
