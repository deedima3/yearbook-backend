package entity

type BlogPage struct {
	PageID      uint64
	HeaderImage string
	Description string
	Owner       uint64
}

type User struct {
	UserID   uint64 `db:"userID"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsActive bool   `db:"isactive"`
	Image    string `db:"image"`
	Nickname string `db:"nickname"`
	Nim      string `db:"nim"`
}

type Users []*User
type BlogPages []*BlogPage