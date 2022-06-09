package helper

import (
	"net/http"
	"os"
	"time"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/user/dto"
	"github.com/deedima3/yearbook-backend/internal/user/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

func JwtDecoder(tokenString string) (string, string) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	BadRequest(err, "Invalid JWT", "Invalid JWT token, please input yang bener")
	var id, nickname string
	for key := range claims {
		if key == "exp" {
			continue
		}
		if key == "id" {
			continue
		}
		if key == "nickname" {
			continue
		}
		sicgolib.NewBaseResponse(400, sicgolib.RESPONSE_ERROR_UNAUTHORIZED_MESSAGE, nil, nil)
	}
	id = claims["id"].(string)
	nickname = claims["nickname"].(string)
	return id, nickname
}

func WrongPass() {
	panic(sicgolib.NewErrorResponse(
		http.StatusBadRequest,
		sicgolib.RESPONSE_ERROR_UNAUTHORIZED_MESSAGE,
		sicgolib.NewErrorResponseValue("Password", "User found but password didn't match"),
	))
}
