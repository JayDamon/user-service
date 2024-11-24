package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var path = "/api/v1/users/{id}/account-tokens"

func TestHandler_CreatePrivateAccessToken_SuccessPathUserExists(t *testing.T) {

	handler := createTestHandler(true)

	userId := uuid.New()

	token := &AccountToken{
		UserId: &userId,
	}

	body, _ := json.Marshal(token)
	req, rec := createRequestAndRecorder("POST", body, userId.String())

	handler.CreatePrivateAccessToken(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, rec.Code, "Expect status created")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Empty(t, data, "Body must be null")
}

func TestHandler_CreatePrivateAccessToken_SuccessPathUserDoesNotExist(t *testing.T) {

	handler := createTestHandler(false)

	userId := uuid.New()

	token := &AccountToken{
		UserId: &userId,
	}

	body, _ := json.Marshal(token)
	req, rec := createRequestAndRecorder("POST", body, userId.String())

	handler.CreatePrivateAccessToken(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, rec.Code, "Expect status created")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Empty(t, data, "Body must be null")
}

func TestHandler_CreatePrivateAccessToken_MissingUserId(t *testing.T) {

	handler := createTestHandler(true)

	token := &AccountToken{}

	body, _ := json.Marshal(token)
	req, rec := createRequestAndRecorder("POST", body, uuid.NewString())

	handler.CreatePrivateAccessToken(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expect status created")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Contains(t, string(data), "valid user id must be provided")
}

func TestHandler_CreatePrivateAccessToken_BodyUserIdDoesNotMatchPath(t *testing.T) {

	handler := createTestHandler(true)

	userId := uuid.New()

	token := &AccountToken{
		UserId: &userId,
	}

	body, _ := json.Marshal(token)
	req, rec := createRequestAndRecorder("POST", body, uuid.NewString())

	handler.CreatePrivateAccessToken(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expect status created")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Contains(t, string(data), "user id from body does not match path variable")
}

func TestHandler_CreatePrivateAccessToken_BodyPathContainsInvalidUUID(t *testing.T) {

	handler := createTestHandler(true)

	userId := uuid.New()

	token := &AccountToken{
		UserId: &userId,
	}

	body, _ := json.Marshal(token)
	req, rec := createRequestAndRecorder("POST", body, "test")

	handler.CreatePrivateAccessToken(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expect status created")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Contains(t, string(data), "path userId variable must be valid uuid")
}

func TestHandler_GetPrivateAccessTokens_HappyPath(t *testing.T) {

	handler := createTestHandler(true)
	testRepo := (handler.UserRepository).(*TestRepository)

	userId := uuid.New()

	accountTokens := make([]*AccountToken, 0)
	token := "testToken"
	itemId := "testItemId"
	cursor := "testCursor"
	accountTokens = append(accountTokens, &AccountToken{UserId: &userId, PrivateToken: &token, ItemID: &itemId, Cursor: &cursor})

	testRepo.mockTokens(accountTokens)

	req, rec := createRequestAndRecorder("POST", nil, userId.String())

	handler.GetPrivateAccessTokens(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, rec.Code, "Expect status OK")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.NotEmpty(t, data, "Body must not be null")

	assert.Equal(t, fmt.Sprintf("[{\"id\":\"%s\",\"privateToken\":\"%s\",\"itemId\":\"%s\",\"cursor\":\"%s\"}]", userId, token, itemId, cursor), string(data))

}

func TestHandler_GetPrivateAccessTokens_UserIdIsNotValidUUID(t *testing.T) {
	handler := createTestHandler(true)

	req, rec := createRequestAndRecorder("POST", nil, "test")

	handler.GetPrivateAccessTokens(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expect status OK")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Contains(t, string(data), "path userId variable must be valid uuid")
}

func TestHandler_GetPrivateAccessTokens_UserDoesNotExist(t *testing.T) {
	handler := createTestHandler(false)

	userId := uuid.New()

	req, rec := createRequestAndRecorder("POST", nil, userId.String())

	handler.GetPrivateAccessTokens(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNotFound, rec.Code, "Expect status OK")

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err, "There should be no error")
	assert.Contains(t, string(data), "user not found")
}

func createRequestAndRecorder(method string, body []byte, id string) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	if body == nil {
		req = httptest.NewRequest(method, path, nil)
	} else {
		req = httptest.NewRequest(method, path, strings.NewReader(string(body)))
	}
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	return req, rec
}

func createTestHandler(userExists bool) *Handler {
	var repository Repository
	if userExists {
		repository = newTestRepository()
	} else {
		repository = newTestRepositoryUserDoesNotExist()
	}
	return NewHandler(repository)
}
