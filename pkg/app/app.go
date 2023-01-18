package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/factotum/moneymaker/user-service/pkg/routes"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Router  *chi.Mux
	Server  *http.Server
	Context *config.Context
}

func (a *App) Initialize(configuration *config.Config) {
	db := connectToDB(configuration)

	a.Context = &config.Context{
		DB:     db,
		Config: configuration,
	}
	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", configuration.HostPort),
		Handler: routes.CreateRoutes(a.Context),
	}
	performDbMigration(db, configuration)
}

func (a *App) Run() {

	defer a.Context.DB.Close()

	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
