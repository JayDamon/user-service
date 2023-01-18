package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httptoolbox "github.com/jaydamon/http-toolbox"
)

func (ctx *UserContext) CreatePrivateAccessToken(w http.ResponseWriter, r *http.Request) {

	// Need to validate that the user id matches the keycloak token and
	// the path variable matches the user token and that they all match the userid from the body

	var usrToken UserAccountToken

	httptoolbox.ReadJsonBodyToVariable(w, r, &usrToken)

	dbCtx := &UserDBContext{
		db: ctx.context.DB,
	}

	usr, err := dbCtx.GetUserById(*usrToken.UserID)
	if err != nil {
		msg := fmt.Sprintf("Unable to create user with id %s", usrToken.UserID)
		log.Println(msg, err)
		httptoolbox.RespondError(w, http.StatusInternalServerError, msg)
	}

	if usr.ID == nil {
		fmt.Println("Creating new user")
		usr.ID = usrToken.UserID
		err = dbCtx.CreateUser(usr)
		if err != nil {
			httptoolbox.RespondError(w, http.StatusInternalServerError, err.Error())
		}
	}

	err = dbCtx.CreateUserAccountToken(&usrToken)
	if err != nil {
		json, _ := json.Marshal(usrToken)
		msg := fmt.Sprintf("Unable to create User Account Token from %s", string(json))
		log.Println(msg, err)
		httptoolbox.RespondError(w, http.StatusInternalServerError, msg)
	}

	httptoolbox.RespondNoBody(w, http.StatusCreated)
}
