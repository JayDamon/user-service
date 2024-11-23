package routes

import (
	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/jaydamon/moneymakergocloak"
	"net/http"

	"github.com/factotum/moneymaker/user-service/pkg/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func CreateRoutes(config *config.Config, handler *user.Handler) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	keyCloakMiddleware := moneymakergocloak.NewMiddleWare(config.KeyCloakConfig)
	router.Use(keyCloakMiddleware.AuthorizeHttpRequest)
	router.Use(middleware.Heartbeat("/ping"))

	user.AddRoutes(router, handler)

	return router
}
