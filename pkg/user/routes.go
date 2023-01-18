package user

import (
	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(mux *chi.Mux, context *config.Context) {

	usrContext := &UserContext{
		context: context,
	}

	mux.Post("/api/v1/users/{id}/account-tokens", usrContext.CreatePrivateAccessToken)

}
