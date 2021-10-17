package persistent

import (
	"context"

	"github.com/EmilGeorgiev/training/go-testing/userservice"
	"github.com/clouway/godb"
	"github.com/google/uuid"
)

type userRepository struct {
	db godb.Database
}

func NewUserRepository(db godb.Database) userservice.UserRepository {
	return userRepository{
		db: db,
	}
}

func (r userRepository) Create(ctx context.Context, nu userservice.NewUser) (userservice.User, error) {
	u := userservice.User{
		ID:                   uuid.New().String(),
		Name:                 nu.Name,
		IdentificationNumber: nu.IdentificationNumber,
	}

	if err := r.db.Collection("users").Insert(u); err != nil {
		return userservice.User{}, err
	}

	return u, nil
}

func (r userRepository) Get(ctx context.Context, id string) (userservice.User, error) {
	var u userservice.User
	if err := r.db.Collection("users").FindID(id).One(&u); err != nil {
		return userservice.User{}, err
	}
	return u, nil
}

func (r userRepository) Update(ctx context.Context, id string, user userservice.NewUser) (userservice.User, error) {
	panic("implement me")
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}