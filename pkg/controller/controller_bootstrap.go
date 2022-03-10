package controller

import (
	"database/sql"

	"github.com/SIC-Unud/sicgolib"
	blogpostControllerPkg "github.com/deedima3/yearbook-backend/internal/blogpost/controller"
	blogpostRepositoryPkg "github.com/deedima3/yearbook-backend/internal/blogpost/repository/impl"
	blogpostServicePkg "github.com/deedima3/yearbook-backend/internal/blogpost/service/impl"
	"github.com/gorilla/mux"
)

func SetupController(router *mux.Router, db *sql.DB) {
	router.Use(sicgolib.ErrorHandlingMiddleware)

	blogpostRouter := router.PathPrefix(API_ROOT_WEB_PATH).Subrouter()
	blogPostRepository := blogpostRepositoryPkg.ProvideBlogpostRepository(db)
	blogPostService := blogpostServicePkg.ProvideRegistrationRepository(blogPostRepository)
	blogPostController := blogpostControllerPkg.ProvideBlogpostController(blogpostRouter, blogPostService)
	blogPostController.InitializeController()
}