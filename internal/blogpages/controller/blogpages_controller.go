package controller

import (
	"net/http"
	"strconv"

	"github.com/SIC-Unud/sicgolib"
	blogpagesServicePkg "github.com/deedima3/yearbook-backend/internal/blogpages/service/api"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/gorilla/mux"
)

type BlogpagesController struct {
	router *mux.Router
	bps blogpagesServicePkg.BlogpagesService
}

func(bpc *BlogpagesController)viewUserPages(rw http.ResponseWriter, r *http.Request){
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

func(bpc *BlogpagesController)InitializeController(){
	bpc.router.HandleFunc(global.API_GET_USER_PAGES, bpc.viewUserPages).Methods(http.MethodGet)
}

func ProvideBlogpagesController(router *mux.Router, bps blogpagesServicePkg.BlogpagesService) *BlogpagesController{
	return &BlogpagesController{router: router, bps: bps}
}