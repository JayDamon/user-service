package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type User struct {
	ID *uuid.UUID `json:"id"`
}

type UserAccountToken struct {
	UserID       *uuid.UUID `json:"id"`
	PrivateToken *string    `json:"privateToken"`
	ItemID       *string    `json:"itemId"`
}

type UserDBContext struct {
	db *sql.DB
}

func (ctx *UserDBContext) GetUserById(id uuid.UUID) (*User, error) {

	db := ctx.db

	statement := `SELECT user_id FROM app_user WHERE user_id = $1`
	log.Printf("Running query: %s\n", statement) // Turn into debug statement
	row := db.QueryRow(statement, id.String())

	var user *User = new(User)
	err := row.Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return userResponse(user)
		}

		log.Println("Error retrieving user", err)
		return nil, err
	}

	return userResponse(user)
}

func userResponse(user *User) (*User, error) {
	json, _ := json.Marshal(user)
	fmt.Printf("User returned \n%s\n", json)

	return user, nil
}

func (ctx *UserDBContext) CreateUser(user *User) error {

	if user.ID == nil {
		return fmt.Errorf("valid Keycloak User id must be provided. Was nil")
	}

	db := ctx.db

	statement := `INSERT INTO app_user (user_id) values ($1)`
	log.Printf("Running query: %s\n", statement) // Turn into debug statement
	_, err := db.Exec(statement, user.ID)
	if err != nil {
		json, _ := json.Marshal(user)
		fmt.Println("Error creating new user", json, err)
		return err
	}

	json, _ := json.Marshal(user)
	fmt.Printf("User created \n%s\n", json)
	fmt.Println(user.ID, &user.ID, *user.ID)

	return nil
}

func (ctx *UserDBContext) CreateUserAccountToken(usrAcctTkn *UserAccountToken) error {

	db := ctx.db

	statement := `INSERT INTO user_account_token (user_id, private_token, item_id) VALUES($1, $2, $3)`
	log.Printf("Running query: %s\nValues: userID = %s, privateToken = %s, itemId = %s\n", statement, usrAcctTkn.UserID.String(), *usrAcctTkn.PrivateToken, *usrAcctTkn.ItemID) // Turn into debug statement
	_, err := db.Exec(statement, usrAcctTkn.UserID, usrAcctTkn.PrivateToken, usrAcctTkn.ItemID)
	if err != nil {
		json, _ := json.Marshal(usrAcctTkn)
		fmt.Println("Error creating new user account token", json, err)
		return err
	}

	return nil
}
