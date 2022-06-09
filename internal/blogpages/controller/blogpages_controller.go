package controller

import (
	"net/http"
	"strconv"

	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
	"github.com/deedima3/yearbook-backend/internal/user/helper"

	"github.com/SIC-Unud/sicgolib"
	blogpagesServicePkg "github.com/deedima3/yearbook-backend/internal/blogpages/service/api"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/gorilla/mux"
)

type BlogpagesController struct {
	router *mux.Router
	bps    blogpagesServicePkg.BlogpagesService
}

func (bpc *BlogpagesController) viewUserPages(rw http.ResponseWriter, r *http.Request) {
	queryVar := mux.Vars(r)
	userID := queryVar["userID"]
	idConv, _ := strconv.ParseUint(userID, 10, 64)

	userPages, err := bpc.bps.ViewUserPages(r.Context(), idConv)
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, userPages).SendResponse(&rw)
}

func (bpc *BlogpagesController) getAllPages(rw http.ResponseWriter, r *http.Request) {
	allPages, err := bpc.bps.GetAllPages(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allPages).SendResponse(&rw)
}

func (u BlogpagesController) NewBlogpage(rw http.ResponseWriter, r *http.Request) {
	bodyNewBlogpage := new(dto.RequestNewBlogpage)
	err := bodyNewBlogpage.FromJSON(r.Body)
	helper.BadRequest(err, "Body format", "Invalid Json format")
	err = u.bps.NewUserPages(r.Context(), *bodyNewBlogpage)
	helper.HelperIfError(err)
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, "success").SendResponse(&rw)
}

func (bpc *BlogpagesController) InitializeController() {
	bpc.router.HandleFunc(global.API_GET_USER_PAGES, bpc.viewUserPages).Methods(http.MethodGet, http.MethodOptions)
	bpc.router.HandleFunc(global.API_GET_ALL_PAGES, bpc.getAllPages).Methods(http.MethodGet, http.MethodOptions)
	bpc.router.HandleFunc(global.API_NEW_BLOGPAGE, bpc.NewBlogpage).Methods(http.MethodPost, http.MethodOptions)
}

func ProvideBlogpagesController(router *mux.Router, bps blogpagesServicePkg.BlogpagesService) *BlogpagesController {
	return &BlogpagesController{router: router, bps: bps}
}
