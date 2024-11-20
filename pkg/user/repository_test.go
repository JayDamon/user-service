package user

import (
	"database/sql"
	"github.com/google/uuid"
)

type TestRepository struct {
	userExists    bool
	accountTokens []*AccountToken
}

func newTestRepository() *TestRepository {
	return &TestRepository{
		userExists: true,
	}
}

func newTestRepositoryUserDoesNotExist() *TestRepository {
	return &TestRepository{
		userExists:    false,
		accountTokens: make([]*AccountToken, 0),
	}
}

func (repository *TestRepository) mockTokens(tokens []*AccountToken) {
	repository.accountTokens = tokens
}

func (repository *TestRepository) GetUserById(id *uuid.UUID) (*User, error) {
	if repository.userExists {
		user := &User{}
		user.ID = id
		return user, nil
	}
	return nil, sql.ErrNoRows
}

func (repository *TestRepository) CreateUser(user *User) error {
	return nil
}

func (repository *TestRepository) GetUserAccountTokensByUserId(userId *uuid.UUID) ([]*AccountToken, error) {
	return repository.accountTokens, nil
}

func (repository *TestRepository) CreateUserAccountToken(accountToken *AccountToken) error {
	return nil
}

func (repository *TestRepository) UpdateUserAccountToken(accountToken *AccountToken) error {
	return nil
}
