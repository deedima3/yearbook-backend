package main

import (
	"os"
	"strings"

	"github.com/SIC-Unud/sicgolib"
	controller "github.com/deedima3/yearbook-backend/pkg/controller"
	"github.com/deedima3/yearbook-backend/pkg/database"
	"github.com/deedima3/yearbook-backend/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initializeGlobalRouter(envVariables map[string]string) *mux.Router {
	r := mux.NewRouter()

	arrayWhiteListedUrls := strings.Split(envVariables["WHITELISTED_URLS"], ",")

	whiteListedUrls := make(map[string]bool)
	
	for _, v := range arrayWhiteListedUrls {
		whiteListedUrls[v] = true
	}

	r.Use(sicgolib.ContentTypeMiddleware)
	r.Use(middleware.CorsMiddleware(whiteListedUrls))
	return r
}

func getEnvVariableValues() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")

	envVariables["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	envVariables["DB_USERNAME"] = os.Getenv("DB_USERNAME")
	envVariables["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVariables["DB_NAME"] = os.Getenv("DB_NAME")

	envVariables["WHITELISTED_URLS"] = os.Getenv("WHITELISTED_URLS")

	return envVariables
}

func main() {
	godotenv.Load()
	envVariables := getEnvVariableValues()
	db := database.GetDatabase()
	r := initializeGlobalRouter(envVariables)

	controller.SetupController(r, db)
	s := sicgolib.ProvideServer(envVariables["SERVER_ADDRESS"], r)
	s.ListenAndServe()
}
