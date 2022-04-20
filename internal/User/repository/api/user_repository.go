package api

import (
	"context"
	"github.com/deedima3/yearbook-backend/internal/User/entity"
)

type UserRepository interface {
	InsertNewUser(ctx context.Context, user entity.User) error
}
