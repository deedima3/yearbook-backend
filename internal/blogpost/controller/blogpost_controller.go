package controller

import (
	blogpostServicePkg "github.com/deedima3/yearbook-backend/internal/blogpost/service/api"
	"github.com/gorilla/mux"
)

type BlogpostController struct {
	router *mux.Router
	bs blogpostServicePkg.BlogpostService
}

func(bc *BlogpostController) InitializeController() {
	//Add your routes here
}

func ProvideBlogpostController(router *mux.Router, bs blogpostServicePkg.BlogpostService) *BlogpostController{
	return &BlogpostController{router: router, bs: bs}
}