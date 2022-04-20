package helper

import (
	"github.com/SIC-Unud/sicgolib"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HelperIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func HelperInternalServerErrorResponse(err error) {
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("create post", "internal server error: "+err.Error())))
	}
}

func BadRequest(err error) {
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("request body", "invalid json format"),
		))
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
