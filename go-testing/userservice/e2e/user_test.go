package e2e

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestWhenCreateANewUser_ThenNewCreateUserIsReturnedAsResponse(t *testing.T) {
	// Arrange
	payload := strings.NewReader(`{"name":"Ivan", "identificationNumber": "951010101496"}`)

	// Act
	resp, _ := http.Post("http://localhost:3000/users", "application/json", payload)
	defer resp.Body.Close()
	var got User
	_ = json.NewDecoder(resp.Body).Decode(&got)

	// Assert
	want := User{
		ID: got.ID,
		Name: "Ivan",
		IdentificationNumber: "951010101496",
	}
	assert.Equal(t, want, got)
}

func TestWhenCreateANewUser_ThenItCanBeGetByID(t *testing.T) {
	// Arrange
	payload := strings.NewReader(`{"name":"Ivan", "identificationNumber": "951212121496"}`)
	resp, _ := http.Post("http://localhost:3000/users", "application/json", payload)
	defer resp.Body.Close()
	var u User
	_ = json.NewDecoder(resp.Body).Decode(&u)

	// Act
	resp2, _ := http.Get("http://localhost:3000/users/" + u.ID)
	defer resp2.Body.Close()
	var got User
	_ = json.NewDecoder(resp2.Body).Decode(&got)

	// Assert
	want := User{
		ID: got.ID,
		Name: "Ivan",
		IdentificationNumber: "951212121496",
	}
	assert.Equal(t, want, got)
}

type User struct {
	ID string
	Name string
	IdentificationNumber string
}