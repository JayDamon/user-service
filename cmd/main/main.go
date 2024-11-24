package main

import (
	"log"

	"github.com/factotum/moneymaker/user-service/pkg/app"
	"github.com/factotum/moneymaker/user-service/pkg/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}

	log.Print("Initializing app\n")

	app.Initialize(config)

	defer app.RabbitConnection.Close()
	app.InitializeRabbitReceivers()

	log.Printf("Starting service on port %s\n", config.HostPort)
	app.Run()
}
