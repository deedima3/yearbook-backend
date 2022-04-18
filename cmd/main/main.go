package main

import (
	"os"

	"github.com/SIC-Unud/sicgolib"
	controller "github.com/deedima3/yearbook-backend/pkg/controller"
	"github.com/deedima3/yearbook-backend/pkg/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initializeGlobalRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(sicgolib.ContentTypeMiddleware)
	return r
}

func getEnvVariableValues() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")
	// envVariables["FIREBASE_CREDENTIALS_PATH"] = os.Getenv("FIREBASE_CREDENTIALS_PATH")

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
	r := initializeGlobalRouter()

	controller.SetupController(r, db)
	s := sicgolib.ProvideServer(envVariables["SERVER_ADDRESS"], r)
	s.ListenAndServe()
}
