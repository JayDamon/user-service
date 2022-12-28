package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/factotum/moneymaker/user-service/pkg/routes"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Router *chi.Mux
	Server *http.Server
	DB     *sql.DB
}

func (a *App) Initialize(config *config.Config) {
	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", config.HostPort),
		Handler: routes.CreateRoutes(a.Broker),
	}
}

func (a *App) Run() {
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
