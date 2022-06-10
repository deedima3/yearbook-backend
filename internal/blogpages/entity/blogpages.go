package entity

type BlogPage struct {
	PageID      uint64 `db:"blogID"`
	HeaderImage string `db:"header_img"`
	Description string `db:"description"`
	Owner       uint64 `db:"owner"`
}

type User struct {
	UserID    uint64 `db:"userID"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	IsActive  bool   `db:"isactive"`
	Image     string `db:"image"`
	Nickname  string `db:"nickname"`
	Nim       string `db:"nim"`
	BirthDate string `db:"birthdate"`
}

type BlogPageUserJoined struct {
	User
	BlogPage
}

type BlogPagesPeopleJoined []*BlogPageUserJoined
type Users []*User
type BlogPages []*BlogPage
