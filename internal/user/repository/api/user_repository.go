package api

import (
	"context"
	"github.com/deedima3/yearbook-backend/internal/user/entity"
)

type UserRepository interface {
	AllUser(ctx context.Context) []entity.User
	InsertNewUser(ctx context.Context, user entity.User) error
	UpdateUser(ctx context.Context, users entity.User) error
}
