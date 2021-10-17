package persistent_test

import (
	"context"
	"testing"

	"github.com/EmilGeorgiev/training/go-testing/userservice"
	"github.com/EmilGeorgiev/training/go-testing/userservice/persistent"
	"github.com/clouway/godb/datastoretest"
	"github.com/stretchr/testify/assert"
)

var db *datastoretest.DB

func TestMain(m *testing.M)  {
	db = datastoretest.NewDatabase()
	m.Run()
	db.Close()
}

func TestUserRepository_Create(t *testing.T) {
	// Arrange
	db.Clean()
	repo := persistent.NewUserRepository(db)
	nu := userservice.NewUser{
		Name:                 "Ivan",
		IdentificationNumber: "950909091496",
	}

	// Act
	u, err1 := repo.Create(context.Background(), nu)
	got, err2 := repo.Get(context.Background(), u.ID)

	// Assert
	want := userservice.User{
		IdentificationNumber: "950909091496",
		ID: got.ID,
		Name: "Ivan",
	}

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, want, got)
	assert.Equal(t, want, u)
}

