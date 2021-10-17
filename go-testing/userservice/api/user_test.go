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
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_WhenUserIsProvided_ThenItIsCreated(t *testing.T) {
	// Arrange
	user := userservice.User{ID: "id", Name: "Ivan", IdentificationNumber: "971010101496"}
	mockRepo := new(MockUserRepository)
	mockRepo.On("Create", userservice.NewUser{Name: "Ivan", IdentificationNumber: "971010101496"}).Return(user, nil)
	h := api.UserHandler{UserRepo: mockRepo}
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
	mockRepo.AssertExpectations(t)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user userservice.NewUser) (userservice.User, error) {
	args := m.Called(user)
	return args.Get(0).(userservice.User), args.Error(1)
}

func (m *MockUserRepository) Get(ctx context.Context, id string) (userservice.User, error) {
	args := m.Called(id)
	return args.Get(0).(userservice.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, id string, user userservice.NewUser) (userservice.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(userservice.User), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
