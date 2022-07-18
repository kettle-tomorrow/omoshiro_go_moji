package controllers

import (
	"encoding/json"
	"net/http"
	"omoshiroGoMoji/backend/models"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	userList, err := json.Marshal(user.List())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userList)
}
