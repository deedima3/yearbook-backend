package helper

import (
	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
	"github.com/deedima3/yearbook-backend/internal/user/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
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

func BadRequest(err error, key string, msg string) {
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue(key, msg),
		))
	}
}

func NotFound(key string, msg string) {
	panic(sicgolib.NewErrorResponse(
		http.StatusNotFound,
		sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
		sicgolib.NewErrorResponseValue(key, msg),
	))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ToUserResponse(user entity.User) dto.UsersResponse {
	return dto.UsersResponse{
		UserID:   user.UserID,
		Email:    user.Email,
		Nickname: user.Nickname,
		Password: user.Password,
	}
}

func JwtTokenGenerate(ID, Nickname string) string {
	claims := jwt.MapClaims{
		"id":       ID,
		"nickname": Nickname,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	HelperIfError(err)
	return t
}

func WrongPass() {
	panic(sicgolib.NewErrorResponse(
		http.StatusBadRequest,
		sicgolib.RESPONSE_ERROR_UNAUTHORIZED_MESSAGE,
		sicgolib.NewErrorResponseValue("Password", "User found but password didn't match"),
	))
}
