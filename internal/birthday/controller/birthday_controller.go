package controller

import (
	"net/http"
	"strconv"

	"github.com/SIC-Unud/sicgolib"
	birthdayServicePkg "github.com/deedima3/yearbook-backend/internal/birthday/service/api"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/gorilla/mux"
)

type BirthdayController struct {
	router *mux.Router
	bs     birthdayServicePkg.BirthdayService
}

func (bc *BirthdayController) CheckUserBirthday(rw http.ResponseWriter, r *http.Request) {
	queryVar := mux.Vars(r)
	userID := queryVar["owner"]
	idConv, _ := strconv.ParseUint(userID, 10, 64)

	isBirthday, err := bc.bs.CheckUserBirthday(r.Context(), idConv)
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, isBirthday).SendResponse(&rw)
}

func (bc *BirthdayController) GetBirthdayWeek(rw http.ResponseWriter, r *http.Request) {
	birthdayWeek, err := bc.bs.GetBirthdayWeek(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("internal", "server error"),
		))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, birthdayWeek).ToJSON(rw)
}

func (bc *BirthdayController) InitializeController() {
	bc.router.HandleFunc(global.API_BIRTHDAY_WEEK, bc.GetBirthdayWeek).Methods(http.MethodGet, http.MethodOptions)
	bc.router.HandleFunc(global.API_IS_BIRTHDAY, bc.CheckUserBirthday).Methods(http.MethodGet, http.MethodOptions)
}

func ProvideBirthdayController(router *mux.Router, bs birthdayServicePkg.BirthdayService) *BirthdayController {
	return &BirthdayController{router: router, bs: bs}
}
