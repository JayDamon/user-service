package user

import (
	"encoding/json"
	"fmt"
	"github.com/jaydamon/moneymakergocloak"
	"github.com/jaydamon/moneymakerrabbit"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type CursorReceiver struct {
	rabbitConnection  moneymakerrabbit.Connector
	goCloakMiddleWare moneymakergocloak.Middleware
	userRepository    Repository
}

type Receiver interface {
	HandleCursorUpdateEvent(msg *amqp091.Delivery) error
}

func NewReceiver(
	rabbitConnection moneymakerrabbit.Connector,
	goCloakMiddleWare moneymakergocloak.Middleware,
	userRepository Repository) Receiver {
	return &CursorReceiver{
		rabbitConnection:  rabbitConnection,
		goCloakMiddleWare: goCloakMiddleWare,
		userRepository:    userRepository,
	}
}

func (receiver CursorReceiver) HandleCursorUpdateEvent(msg *amqp091.Delivery) error {

	log.Println("Received Message from update_cursor queue")

	err := receiver.goCloakMiddleWare.AuthorizeMessage(msg)
	if err != nil {
		fmt.Printf("unauthorized message. %s\n", err)
		return err
	}
	token, err := moneymakergocloak.GetAuthorizationHeaderFromMessage(msg)
	if err != nil {
		fmt.Printf("error when extracting token from request. %s\n", err)
		return err
	}
	userId, err := receiver.goCloakMiddleWare.ExtractUserIdFromToken(&token)
	if err != nil {
		fmt.Printf("error extracting user id from jwt token. %s\n", err)
		return err
	}
	fmt.Println("successfully authorized message")

	log.Printf("Processing body %s\n", msg.Body)
	var at AccountToken
	err = json.Unmarshal(msg.Body, &at)
	if err != nil {
		log.Printf("Unable to unmarshal body to Private Token object \n%s\n", msg.Body)
		return err
	}
	log.Printf("Unmarshalled message body to Private Token object %+v\n", at)

	if userId != (*at.UserId).String() {
		log.Printf("invalid private token. user id does not match oauth token")
		return err
	}

	err = receiver.userRepository.UpdateUserAccountToken(&at)
	if err != nil {
		log.Printf("Unable to save cursor updates \n%s\n", err)
		return err
	}
	return nil
}
