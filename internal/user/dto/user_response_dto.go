package dto

type UsersResponse struct {
	UserID   uint64 `json:"userID"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
