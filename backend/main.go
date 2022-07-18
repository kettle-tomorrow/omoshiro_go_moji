package main

import (
	"context"
	"net/http"
	"omoshiroGoMoji/backend/controllers"
	"omoshiroGoMoji/backend/databases"
	"omoshiroGoMoji/backend/models"

	"github.com/go-chi/chi/v5"
)

func main() {
	databases.Init(&models.User{}, &models.OmoshiroGoMoji{}, &models.Account{})
	defer databases.Close()

	router := chi.NewRouter()

	router.Get("/api/v1/omoshiro_go_moji/list", controllers.OmoshiroGoMojiIndex)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		apiRouter.Post("/login", controllers.AccountSessionCreate)
		apiRouter.Post("/account", controllers.AccountCreate)
		apiRouter.Route("/omoshiro_go_moji", func(omoshiroGoMojiRouter chi.Router) {
			omoshiroGoMojiRouter.Get("/list", controllers.OmoshiroGoMojiIndex)
			omoshiroGoMojiRouter.Get("/{id}", controllers.OmoshiroGoMojiShow)
			omoshiroGoMojiRouter.Use(loginCheckMiddleware)
			omoshiroGoMojiRouter.Post("", controllers.OmoshiroGoMojiCreate)
			omoshiroGoMojiRouter.Patch("/{id}", controllers.OmoshiroGoMojiUpdate)
			omoshiroGoMojiRouter.Delete("/{id}", controllers.OmoshiroGoMojiDelete)
		})
		apiRouter.Get("/user/list", controllers.UserIndex)
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	http.ListenAndServe(":3000", router)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, PUT, OPTIONS")
		next.ServeHTTP(w, r)
	})
}

func loginCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_cookie")
		if err == nil {
			ctx := context.WithValue(r.Context(), "currentUserId", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
