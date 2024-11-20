package app

import (
	"github.com/factotum/moneymaker/user-service/pkg/user"
	"github.com/jaydamon/moneymakergocloak"
)

func (a *App) InitializeRabbitReceivers() {

	goCloakMiddleWare := moneymakergocloak.NewMiddleWare(a.Config.KeyCloakConfig)

	go a.RabbitConnection.ReceiveMessages(
		"update_cursor",
		user.NewReceiver(
			a.RabbitConnection,
			goCloakMiddleWare,
			a.UserRepository).HandleCursorUpdateEvent)
}
