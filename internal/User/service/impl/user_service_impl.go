package impl

import (
	"context"
	"github.com/deedima3/yearbook-backend/internal/User/dto"
	"github.com/deedima3/yearbook-backend/internal/User/entity"
	"github.com/deedima3/yearbook-backend/internal/User/helper"
	UserRepositoryPkg "github.com/deedima3/yearbook-backend/internal/User/repository/api"
)

type userServiceImpl struct {
	ur UserRepositoryPkg.UserRepository
}

func ProvideUserService(repository UserRepositoryPkg.UserRepository) *userServiceImpl {
	return &userServiceImpl{
		ur: repository,
	}
}

func (u userServiceImpl) CreateUser(ctx context.Context, body dto.UserRegisterRequestBody) error {
	hashres, err := helper.HashPassword(body.Password)
	helper.HelperIfError(err)
	err = u.ur.InsertNewUser(ctx, entity.User{
		Email:    body.Email,
		Password: hashres,
		Nickname: body.Nickname,
		Nim:      body.Nim,
	})
	helper.HelperInternalServerErrorResponse(err)
	return nil
}
