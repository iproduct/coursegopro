package persistent_test

import (
	"context"
	"testing"

	"github.com/EmilGeorgiev/training/go-testing/userservice"
	"github.com/EmilGeorgiev/training/go-testing/userservice/persistent"
	"github.com/clouway/godb"
	"github.com/stretchr/testify/assert"
)


func TestUserRepository_CreateWithFakeDB(t *testing.T) {
	// Arrange
	fDB := NewFakeDB()
	repo := persistent.NewUserRepository(fDB)
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

type fakeDB struct {
	collection godb.Collection
	godb.Database
}

func NewFakeDB() godb.Database {
	return fakeDB{
		collection: inmemoryCollection{users: map[string]userservice.User{}},
		Database: nil,
	}
}

func (i fakeDB) Close() {}

func (i fakeDB) Collection(name string) godb.Collection {
	return i.collection
}

type inmemoryCollection struct {
	users map[string]userservice.User
	godb.Collection
}

func (i inmemoryCollection) FindID(id interface{}) godb.Query {
	u := i.users[id.(string)]
	return fakeQuery{user: u, Query:nil}
}

func (i inmemoryCollection) Insert(doc interface{}) error {
	u := doc.(userservice.User)
	i.users[u.ID] = u
	return nil
}

type fakeQuery struct {
	user userservice.User
	godb.Query
}

func (f fakeQuery) One(result interface{}) error {
	u := result.(*userservice.User)
	*u = f.user
	return nil
}
