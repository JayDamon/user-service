package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type PostgresRepository struct {
	Conn *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}

func (repository *PostgresRepository) GetUserById(id *uuid.UUID) (*User, error) {

	db := repository.Conn

	statement := `SELECT user_id FROM app_user WHERE user_id = $1`
	log.Printf("Running query: %s\n", statement) // Turn into debug statement
	row := db.QueryRow(statement, id.String())

	var user = new(User)
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

func (repository *PostgresRepository) CreateUser(user *User) error {

	if user.ID == nil {
		return fmt.Errorf("valid Keycloak User id must be provided. Was nil")
	}

	db := repository.Conn

	statement := `INSERT INTO app_user (user_id) values ($1)`
	log.Printf("Running query: %s\n", statement) // Turn into debug statement
	_, err := db.Exec(statement, user.ID)
	if err != nil {
		jsonUser, _ := json.Marshal(user)
		fmt.Println("Error creating new user", jsonUser, err)
		return err
	}

	jsonUser, _ := json.Marshal(user)
	fmt.Printf("User created \n%s\n", jsonUser)
	fmt.Println(user.ID, &user.ID, *user.ID)

	return nil
}

func (repository *PostgresRepository) GetUserAccountTokensByUserId(userId *uuid.UUID) ([]*AccountToken, error) {

	db := repository.Conn
	statement := `SELECT user_id, private_token, item_id, plaid_cursor FROM user_account_token WHERE user_id = $1`

	log.Printf("Running query: %s\nValues: userId = %s\n", statement, userId)

	results, err := db.Query(statement, userId)
	if err != nil {
		log.Printf("Error retrieving user account token for user id %s causing error %s\n", userId, err)
		return nil, err
	}

	tokens := make([]*AccountToken, 0)
	for results.Next() {
		var token AccountToken
		err := results.Scan(&token.UserId, &token.PrivateToken, &token.ItemID, &token.Cursor)
		if err != nil {
			log.Printf("error caused while scanning results of query\nQuery: %s\nError: %s\n", statement, err)
			return nil, err
		}
		tokens = append(tokens, &token)
	}

	return tokens, nil
}

func (repository *PostgresRepository) CreateUserAccountToken(accountToken *AccountToken) error {

	db := repository.Conn

	statement := `INSERT INTO user_account_token (user_id, private_token, item_id) VALUES($1, $2, $3)`
	log.Printf("Running query: %s\nValues: userID = %s, privateToken = %s, itemId = %s\n", statement, accountToken.UserId.String(), *accountToken.PrivateToken, *accountToken.ItemID) // Turn into debug statement
	_, err := db.Exec(statement, accountToken.UserId, accountToken.PrivateToken, accountToken.ItemID)
	if err != nil {
		jsonAccountToken, _ := json.Marshal(accountToken)
		fmt.Println("Error creating new user account token", jsonAccountToken, err)
		return err
	}

	return nil
}

func (repository *PostgresRepository) UpdateUserAccountToken(accountToken *AccountToken) error {

	db := repository.Conn
	statement := `UPDATE user_account_token SET plaid_cursor = $1 where user_id = $2 and item_id = $3`
	log.Printf("Running query: %s\nValues: cursor = %s, userID = %s, itemId = %s\n", statement, *accountToken.Cursor, accountToken.UserId.String(), *accountToken.ItemID) // Turn into debug statement
	_, err := db.Exec(statement, accountToken.Cursor, accountToken.UserId, accountToken.ItemID)
	if err != nil {
		jsonAccountToken, _ := json.Marshal(accountToken)
		fmt.Println("Error updating cursor for account token", jsonAccountToken, err)
		return err
	}

	return nil
}

func userResponse(user *User) (*User, error) {
	jsonUser, _ := json.Marshal(user)
	fmt.Printf("User returned \n%s\n", jsonUser)

	return user, nil
}
