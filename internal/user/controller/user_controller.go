package controller

import (
	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
	"github.com/deedima3/yearbook-backend/internal/user/helper"
	userServicePkg "github.com/deedima3/yearbook-backend/internal/user/service/api"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	router *mux.Router
	us     userServicePkg.UserServiceInterface
}

func (u *UserController) CreateUser(rw http.ResponseWriter, r *http.Request) {
	userRequest := new(dto.UserRegisterRequestBody)
	err := userRequest.FromJSON(r.Body)
	helper.BadRequest(err, "body format", "invalid Json format")

	_ = u.us.CreateUser(r.Context(), *userRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, "success").ToJSON(rw)
}

func (u *UserController) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	userUpdate := new(dto.UserUpdateRequestBody)
	err := userUpdate.FromJSON(r.Body)
	helper.BadRequest(err, "body format", "invalid Json format")

	_ = u.us.SaveUser(r.Context(), *userUpdate)

	tokenJWT := struct {
		Accesstoken string `json:"accesstoken"`
	}{
		Accesstoken: helper.JwtTokenGenerate(strconv.FormatUint(userUpdate.UserID, 10), userUpdate.Nickname),
	}

	sicgolib.NewBaseResponse(204, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, tokenJWT).ToJSON(rw)
}

func (u UserController) AllUser(rw http.ResponseWriter, r *http.Request) {
	users := u.us.GetAllUser(r.Context())
	sicgolib.NewBaseResponse(204, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, users).ToJSON(rw)
}

func (uc *UserController) InitializeController() {
	//Add your routes here
	uc.router.HandleFunc(global.API_INSERT_USER, uc.CreateUser).Methods(http.MethodPost)
	uc.router.HandleFunc(global.API_UPDATE_USER, uc.UpdateUser).Methods(http.MethodPost)
	uc.router.HandleFunc(global.API_ALL_USER, uc.AllUser).Methods(http.MethodGet)
}

func ProvideUserController(router *mux.Router, us userServicePkg.UserServiceInterface) *UserController {
	return &UserController{router: router, us: us}
}
