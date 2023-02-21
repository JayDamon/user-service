package user

import "github.com/google/uuid"

type Repository interface {
	GetUserById(id *uuid.UUID) (*User, error)
	CreateUser(user *User) error
	GetUserAccountTokensByUserId(userId *uuid.UUID) ([]*AccountToken, error)
	CreateUserAccountToken(accountToken *AccountToken) error
}
