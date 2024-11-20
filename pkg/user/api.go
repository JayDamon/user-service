package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"

	httptoolbox "github.com/jaydamon/http-toolbox"
)

type Handler struct {
	UserRepository Repository
}

func NewHandler(userRepository Repository) *Handler {
	return &Handler{
		UserRepository: userRepository,
	}
}

func (handler *Handler) CreatePrivateAccessToken(w http.ResponseWriter, r *http.Request) {

	var usrToken AccountToken

	httptoolbox.ReadJsonBodyToVariable(w, r, &usrToken)

	if usrToken.UserId == nil {
		msg := "valid user id must be provided"
		log.Println(msg)
		httptoolbox.RespondError(w, http.StatusBadRequest, msg)
		return
	}

	userId := chi.URLParam(r, "id")

	pathUuid, err := uuid.Parse(userId)
	if err != nil {
		msg := "path userId variable must be valid uuid"
		log.Println(msg, fmt.Sprintf("%s is not a valid uuid", userId))
		httptoolbox.RespondError(w, http.StatusBadRequest, msg)
		return
	}

	if pathUuid != *usrToken.UserId {
		msg := "user id from body does not match path variable"
		log.Println(msg, fmt.Sprintf("body was '%s' and path was '%s'", usrToken.UserId, userId))
		httptoolbox.RespondError(w, http.StatusBadRequest, msg)
		return
	}

	repository := handler.UserRepository
	usr, err := repository.GetUserById(usrToken.UserId)
	if err != nil {
		if err != sql.ErrNoRows {
			msg := fmt.Sprintf("Unable to create user with id %s", usrToken.UserId)
			log.Println(msg, err)
			httptoolbox.RespondError(w, http.StatusInternalServerError, msg)
			return
		}
		usr = &User{}
	}

	if usr.ID == nil {
		fmt.Println("Creating new user")
		usr.ID = usrToken.UserId
		err = repository.CreateUser(usr)
		if err != nil {
			httptoolbox.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	err = repository.CreateUserAccountToken(&usrToken)
	if err != nil {
		jsonToken, _ := json.Marshal(usrToken)
		msg := fmt.Sprintf("Unable to create User Account Token from %s", string(jsonToken))
		log.Println(msg, err)
		httptoolbox.RespondError(w, http.StatusInternalServerError, msg)
		return
	}

	httptoolbox.RespondNoBody(w, http.StatusCreated)
}

func (handler *Handler) GetPrivateAccessTokens(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "id")
	id, err := uuid.Parse(userId)
	if err != nil {
		msg := "path userId variable must be valid uuid"
		log.Println(msg, fmt.Sprintf("%s is not a valid uuid", userId))
		httptoolbox.RespondError(w, http.StatusBadRequest, msg)
		return
	}

	repository := handler.UserRepository

	user, err := repository.GetUserById(&id)
	if err != nil {
		msg := "user not found"
		log.Println(msg, fmt.Sprintf("%s does not exist in the database", userId))
		httptoolbox.RespondError(w, http.StatusNotFound, msg)
		return
	}

	tokens, err := repository.GetUserAccountTokensByUserId(user.ID)
	if err != nil {
		msg := "Unable to retrieve user access tokens"
		log.Println(msg, fmt.Sprintf("Issue occured retrieving tokens for user with id %s", userId), err)
		httptoolbox.RespondError(w, http.StatusInternalServerError, msg)
		return
	}

	httptoolbox.Respond(w, http.StatusOK, tokens)
}
