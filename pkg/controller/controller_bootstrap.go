package controller

import (
	"database/sql"

	"github.com/SIC-Unud/sicgolib"
	blogpostControllerPkg "github.com/deedima3/yearbook-backend/internal/blogpost/controller"
	blogpostRepositoryPkg "github.com/deedima3/yearbook-backend/internal/blogpost/repository/impl"
	blogpostServicePkg "github.com/deedima3/yearbook-backend/internal/blogpost/service/impl"
	"github.com/deedima3/yearbook-backend/internal/ping"
	"github.com/gorilla/mux"
)

func SetupController(router *mux.Router, db *sql.DB) {
	router.Use(sicgolib.ErrorHandlingMiddleware)

	webRouter := router.NewRoute().Subrouter()

	blogpostRouter := router.PathPrefix(API_ROOT_WEB_PATH).Subrouter()
	blogPostRepository := blogpostRepositoryPkg.ProvideBlogpostRepository(db)
	blogPostService := blogpostServicePkg.ProvideRegistrationRepository(blogPostRepository)
	blogPostController := blogpostControllerPkg.ProvideBlogpostController(blogpostRouter, blogPostService)
	blogPostController.InitializeController()

	pingService := ping.ProvidePingService()
	pingController := ping.ProvidePingController(webRouter, pingService)
	pingController.InitializeController()
}