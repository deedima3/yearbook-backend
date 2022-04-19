package controller

import (
	"net/http"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/blogpost/dto"
	blogpostServicePkg "github.com/deedima3/yearbook-backend/internal/blogpost/service/api"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/gorilla/mux"
)

type BlogpostController struct {
	router *mux.Router
	bs blogpostServicePkg.BlogpostService
}

func(bc *BlogpostController) createPost(rw http.ResponseWriter, r *http.Request){
	postRequest := new(dto.BlogPostRequestBody)

	if err := postRequest.FromJSON(r.Body); err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("request body", "invalid json format"),
		))
	}

	blogID, _ := bc.bs.CreatePost(r.Context(), *postRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, blogID).ToJSON(rw)
}

func(bc *BlogpostController) InitializeController() {
	//Add your routes here
	bc.router.HandleFunc(global.API_INSERT_POST, bc.createPost).Methods(http.MethodPost)
}

func ProvideBlogpostController(router *mux.Router, bs blogpostServicePkg.BlogpostService) *BlogpostController{
	return &BlogpostController{router: router, bs: bs}
}