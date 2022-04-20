package controller

import (
	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
	"github.com/deedima3/yearbook-backend/internal/user/helper"
	userServicePkg "github.com/deedima3/yearbook-backend/internal/user/service/api"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	router *mux.Router
	us     userServicePkg.UserServiceInterface
}

func (u *UserController) CreateUser(rw http.ResponseWriter, r *http.Request) {
	userRequest := new(dto.UserRegisterRequestBody)
	err := userRequest.FromJSON(r.Body)
	helper.BadRequest(err)

	_ = u.us.CreateUser(r.Context(), *userRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, "success").ToJSON(rw)
}

func (uc *UserController) InitializeController() {
	//Add your routes here
	uc.router.HandleFunc(global.API_INSERT_USER, uc.CreateUser).Methods(http.MethodPost)
}

func ProvideUserController(router *mux.Router, us userServicePkg.UserServiceInterface) *UserController {
	return &UserController{router: router, us: us}
}
