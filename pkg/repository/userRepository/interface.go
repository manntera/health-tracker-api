package userRepository

import "context"

type IUserRepository interface {
	AddData(ctx context.Context, userData *User) error
	GetData(ctx context.Context, id string) (*User, error)
	DeleteData(ctx context.Context, id string) error
}
