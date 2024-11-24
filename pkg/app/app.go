package app

import (
	"database/sql"
	"fmt"
	"github.com/factotum/moneymaker/user-service/pkg/user"
	"github.com/jaydamon/moneymakerrabbit"
	"log"
	"net/http"

	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/factotum/moneymaker/user-service/pkg/routes"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Router           *chi.Mux
	Server           *http.Server
	DB               *sql.DB
	Config           *config.Config
	UserRepository   user.Repository
	UserHandler      *user.Handler
	RabbitConnection moneymakerrabbit.Connector
}

func (a *App) Initialize(configuration *config.Config) {
	a.DB = connectToDB(configuration)
	a.Config = configuration
	a.UserRepository = user.NewPostgresRepository(a.DB)
	a.UserHandler = user.NewHandler(a.UserRepository)
	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", configuration.HostPort),
		Handler: routes.CreateRoutes(configuration, a.UserHandler),
	}
	performDbMigration(a.DB, configuration)
	a.RabbitConnection = a.Config.Rabbit.Connect()
}

func (a *App) Run() {

	defer a.DB.Close()

	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
