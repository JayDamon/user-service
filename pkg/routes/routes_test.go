package routes

import (
	"fmt"
	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/factotum/moneymaker/user-service/pkg/user"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateRoutes_RoutesExist(t *testing.T) {

	userApiHandler := user.Handler{}
	configuration := config.Config{}
	routes := CreateRoutes(&configuration, &userApiHandler, false)
	chiRoutes := routes.(chi.Router)

	assert.NotNil(t, chiRoutes)

	routeExists(t, chiRoutes, "/v1/users/{id}/account-tokens", "POST")
	routeExists(t, chiRoutes, "/v1/users/{id}/account-tokens", "GET")
}

func routeExists(t *testing.T, routes chi.Router, routeToValidate string, methodToValidate string) {
	found := false

	_ = chi.Walk(routes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Println(method)
		if route == routeToValidate && method == methodToValidate {
			found = true
		}
		return nil
	})
	assert.True(t, found, "route not found %s", routeToValidate)
}
