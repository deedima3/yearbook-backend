package impl

import (
	"context"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
	"github.com/deedima3/yearbook-backend/internal/user/entity"
	"github.com/deedima3/yearbook-backend/internal/user/helper"
	UserRepositoryPkg "github.com/deedima3/yearbook-backend/internal/user/repository/api"
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

func (u userServiceImpl) SaveUser(ctx context.Context, body dto.UserUpdateRequestBody) error {
	hashres, err := helper.HashPassword(body.Password)
	helper.HelperIfError(err)
	user := entity.User{UserID: body.UserID, Email: body.Email, Password: hashres, Image: body.Email, Nickname: body.Nickname}
	err = u.ur.UpdateUser(ctx, user)
	helper.HelperIfError(err)
	return nil
}

func (u userServiceImpl) GetAllUser(ctx context.Context) []dto.UsersResponse {
	users := u.ur.AllUser(ctx)
	var usersResponses []dto.UsersResponse
	for _, user := range users {
		usersResponses = append(usersResponses, helper.ToUserResponse(user))
	}
	return usersResponses
}