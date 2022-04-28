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
	helper.BadRequest(err, "body format", "Invalid Json format")

	_ = u.us.CreateUser(r.Context(), *userRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, "success").ToJSON(rw)
}

func (u *UserController) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	userUpdate := new(dto.UserUpdateRequestBody)
	err := userUpdate.FromJSON(r.Body)
	helper.BadRequest(err, "body format", "Invalid Json format")

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

func (u UserController) Login(rw http.ResponseWriter, r *http.Request) {
	loginBody := new(dto.LoginRequestBody)
	err := loginBody.FromJSON(r.Body)
	helper.BadRequest(err, "Body format", "Invalid Json format")
	id, nickname := u.us.PassForLogin(r.Context(), *loginBody)
	tokenJWT := struct {
		Accesstoken string `json:"accesstoken"`
	}{
		Accesstoken: helper.JwtTokenGenerate(strconv.FormatUint(id, 10), nickname),
	}

	if id != 0 {
		sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, tokenJWT).ToJSON(rw)
	} else {
		helper.WrongPass()
	}
}

func (u UserController) RefreshToken(rw http.ResponseWriter, r *http.Request) {
	var id, nickname string
	if r.Header.Get("Token") == "" {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("Wrong Key", "Key must be named 'Token' "),
		))
	}
	id, nickname = helper.JwtDecoder(r.Header.Get("Token"))
	newToken := helper.JwtTokenGenerate(id, nickname)
	tokenJWT := struct {
		Accesstoken string `json:"accesstoken"`
	}{
		Accesstoken: newToken,
	}
	rw.Header().Add("Token", newToken)
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, tokenJWT).ToJSON(rw)
}

func (uc *UserController) InitializeController() {
	//Add your routes here
	uc.router.HandleFunc(global.API_INSERT_USER, uc.CreateUser).Methods(http.MethodPost)
	uc.router.HandleFunc(global.API_UPDATE_USER, uc.UpdateUser).Methods(http.MethodPost)
	uc.router.HandleFunc(global.API_ALL_USER, uc.AllUser).Methods(http.MethodGet)
	uc.router.HandleFunc(global.API_LOGIN, uc.Login).Methods(http.MethodPost)
	uc.router.HandleFunc(global.API_REFRESH_TOKEN, uc.RefreshToken).Methods(http.MethodPost)
}

func ProvideUserController(router *mux.Router, us userServicePkg.UserServiceInterface) *UserController {
	return &UserController{router: router, us: us}
}
