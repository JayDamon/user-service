package user

import (
	"database/sql"
	"github.com/google/uuid"
)

type testRepository struct {
	userExists    bool
	accountTokens []*AccountToken
}

func newTestRepository() *testRepository {
	return &testRepository{
		userExists: true,
	}
}

func newTestRepositoryUserDoesNotExist() *testRepository {
	return &testRepository{
		userExists:    false,
		accountTokens: make([]*AccountToken, 0),
	}
}

func (repository *testRepository) mockTokens(tokens []*AccountToken) {
	repository.accountTokens = tokens
}

func (repository *testRepository) GetUserById(id *uuid.UUID) (*User, error) {
	if repository.userExists {
		user := &User{}
		user.ID = id
		return user, nil
	}
	return nil, sql.ErrNoRows
}

func (repository *testRepository) CreateUser(user *User) error {
	return nil
}

func (repository *testRepository) GetUserAccountTokensByUserId(userId *uuid.UUID) ([]*AccountToken, error) {
	return repository.accountTokens, nil
}

func (repository *testRepository) CreateUserAccountToken(accountToken *AccountToken) error {
	return nil
}
