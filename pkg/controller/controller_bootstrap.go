package controller

import (
	"database/sql"

	"github.com/SIC-Unud/sicgolib"
	blogpostControllerPkg "github.com/deedima3/yearbook-backend/internal/blogpost/controller"
	blogpostRepositoryPkg "github.com/deedima3/yearbook-backend/internal/blogpost/repository/impl"
	blogpostServicePkg "github.com/deedima3/yearbook-backend/internal/blogpost/service/impl"
	"github.com/deedima3/yearbook-backend/internal/ping"
	userControllerPkg "github.com/deedima3/yearbook-backend/internal/user/controller"
	userRepositoryPkg "github.com/deedima3/yearbook-backend/internal/user/repository/impl"
	userServicePkg "github.com/deedima3/yearbook-backend/internal/user/service/impl"
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

	userRouter := router.PathPrefix(API_ROOT_WEB_PATH).Subrouter()
	userRepository := userRepositoryPkg.ProvideUserRepository(db)
	userService := userServicePkg.ProvideUserService(userRepository)
	userController := userControllerPkg.ProvideUserController(userRouter, userService)
	userController.InitializeController()

	pingService := ping.ProvidePingService()
	pingController := ping.ProvidePingController(webRouter, pingService)
	pingController.InitializeController()
}
