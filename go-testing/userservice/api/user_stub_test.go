package api_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/EmilGeorgiev/training/go-testing/userservice"
	"github.com/EmilGeorgiev/training/go-testing/userservice/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_WithStubObject(t *testing.T) {
	// Arrange
	user := userservice.User{ID: "id", Name: "Ivan", IdentificationNumber: "971010101496"}
	stubRepo := StubUserRepository{
		user: user,
		err:  nil,
	}
	h := api.UserHandler{UserRepo: stubRepo}
	w := httptest.NewRecorder()

	// Act
	payload := `{"Name":"Ivan", "IdentificationNumber":"971010101496"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	h.Create(w, req)

	// Assert
	var got userservice.User
	_ = json.NewDecoder(w.Body).Decode(&got)

	assert.Equal(t, user, got)
	assert.EqualValues(t, http.StatusCreated, w.Code)
}

type StubUserRepository struct {
	user userservice.User
	err error
}

func (s StubUserRepository) Create(ctx context.Context, user userservice.NewUser) (userservice.User, error) {
	return s.user, s.err
}

func (s StubUserRepository) Get(ctx context.Context, id string) (userservice.User, error) {
	return s.user, s.err
}

func (s StubUserRepository) Update(ctx context.Context, id string, user userservice.NewUser) (userservice.User, error) {
	return s.user, s.err
}

func (s StubUserRepository) Delete(ctx context.Context, id string) error {
	return s.err
}