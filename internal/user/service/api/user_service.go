package api

import (
	"context"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
)

type UserServiceInterface interface {
	GetAllUser(ctx context.Context) []dto.UsersResponse
	CreateUser(ctx context.Context, body dto.UserRegisterRequestBody) error
	SaveUser(ctx context.Context, body dto.UserUpdateRequestBody) error
}
