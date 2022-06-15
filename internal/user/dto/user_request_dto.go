package dto

import (
	"encoding/json"
	"io"
	"time"
)

type UserRegisterRequestBody struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	Nim       string    `json:"nim"`
	BirthDate time.Time `json:"birthDate"`
}

type UserUpdateRequestBody struct {
	UserID   uint64 `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Nim      string `json:"nim"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *LoginRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}

func (u *UserRegisterRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}

func (u *UserUpdateRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}
