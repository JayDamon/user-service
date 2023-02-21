package user

import (
	"github.com/go-chi/chi/v5"
)

func AddRoutes(mux *chi.Mux, controller *Handler) {

	mux.Post("/v1/users/{id}/account-tokens", controller.CreatePrivateAccessToken)
	mux.Get("/v1/users/{id}/account-tokens", controller.GetPrivateAccessTokens)

}
