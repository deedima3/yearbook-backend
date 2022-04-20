package api

import (
	"context"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, body dto.UserRegisterRequestBody) error
}
