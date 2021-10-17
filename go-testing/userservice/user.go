package userservice

import "context"

type UserRepository interface {
	Create(context.Context, NewUser) (User, error)
	Get(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, id string, user NewUser) (User, error)
	Delete(ctx context.Context, id string) error
}

type NewUser struct {
	Name                 string
	IdentificationNumber string
}

type User struct {
	ID                   string `bson:"_id"`
	Name                 string `bson:"name"`
	IdentificationNumber string `bson:"identification_number"`
}