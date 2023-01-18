package routes

import (
	"net/http"

	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/factotum/moneymaker/user-service/pkg/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func CreateRoutes(context *config.Context) http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	addRoutes(mux, context)

	return mux
}

func addRoutes(mux *chi.Mux, context *config.Context) {
	user.AddRoutes(mux, context)
}
