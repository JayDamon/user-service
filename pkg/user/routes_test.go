package user

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAddRoutes(t *testing.T) {

	userApiHandler := Handler{}

	router := chi.NewRouter()

	AddRoutes(router, &userApiHandler)

	assert.NotNil(t, router)

	routeExists(t, router, "/v1/users/{id}/account-tokens", "POST")
	routeExists(t, router, "/v1/users/{id}/account-tokens", "GET")
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
