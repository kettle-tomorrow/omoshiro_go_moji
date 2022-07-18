package controllers

import (
	"encoding/json"
	"net/http"
	"omoshiroGoMoji/backend/models"

	"github.com/go-chi/chi/v5"
)

func OmoshiroGoMojiCreate(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(uint)
	omoshiroGoMoji := models.OmoshiroGoMoji{UserID: userId, Name: r.PostFormValue("name")}
	omoshiroGoMoji.Create(&omoshiroGoMoji)
	omoshiroGoMojiResponse, err := json.Marshal(omoshiroGoMoji)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(omoshiroGoMojiResponse)
}

func OmoshiroGoMojiIndex(w http.ResponseWriter, r *http.Request) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	omoshiroGoMojiList, err := json.Marshal(omoshiroGoMoji.List())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(omoshiroGoMojiList)
}

func OmoshiroGoMojiShow(w http.ResponseWriter, r *http.Request) {
	omoshiroGomoji := models.OmoshiroGoMoji{}
	omoshiroGoMoji := omoshiroGomoji.Get(chi.URLParam(r, "id"))
	if omoshiroGoMoji.ID == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	omoshiroGomojiResponse, err := json.Marshal(omoshiroGoMoji)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(omoshiroGomojiResponse)
}

func OmoshiroGoMojiUpdate(w http.ResponseWriter, r *http.Request) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	omoshiroGoMoji = omoshiroGoMoji.Get(chi.URLParam(r, "id"))
	omoshiroGoMoji.Name = r.PostFormValue("name")
	if omoshiroGoMoji.UserID != r.Context().Value("userId").(uint) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	omoshiroGoMoji.Update(&omoshiroGoMoji)
	omoshiroGomojiResponse, err := json.Marshal(omoshiroGoMoji)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(omoshiroGomojiResponse)
}

func OmoshiroGoMojiDelete(w http.ResponseWriter, r *http.Request) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	omoshiroGoMoji = omoshiroGoMoji.Get(chi.URLParam(r, "id"))
	if omoshiroGoMoji.UserID != r.Context().Value("userId").(uint) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	omoshiroGoMoji.Delete(chi.URLParam(r, "id"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
