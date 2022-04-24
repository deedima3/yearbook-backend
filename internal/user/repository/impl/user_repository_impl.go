package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/deedima3/yearbook-backend/internal/user/entity"
	"github.com/deedima3/yearbook-backend/internal/user/helper"
	"strconv"
	"strings"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func ProvideUserRepository(db *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{DB: db}
}

func (u userRepositoryImpl) InsertNewUser(ctx context.Context, user entity.User) error {
	SQLQUERY := "INSERT INTO yearbook_db.user(email,password,nickname,nim) VALUES(?,?,?,?)"
	stmt, err := u.DB.PrepareContext(ctx, SQLQUERY)
	helper.HelperIfError(err)
	res, err := stmt.ExecContext(ctx, strings.ToLower(user.Email), user.Password, user.Nickname, user.Nim)
	helper.HelperIfError(err)
	_, err = res.LastInsertId()
	helper.HelperIfError(err)
	return nil
}

func (u userRepositoryImpl) UpdateUser(ctx context.Context, users entity.User) error {
	SQLQUERY := "SELECT userID FROM yearbook_db.user WHERE userID = ?"
	rows, err := u.DB.QueryContext(ctx, SQLQUERY, users.UserID)
	helper.HelperIfError(err)
	if rows.Next() {
		err := rows.Scan(&users.UserID)
		helper.HelperIfError(err)
	} else {
		fmt.Println("TIDAK DITEMUKAN")
		helper.NotFound("userID "+strconv.FormatUint(users.UserID, 10)+" not found", "Masukan ID user yang ada")
	}

	SQLQUERY = "UPDATE yearbook_db.user SET email = ?, password = ?, image = ?, nickname = ? WHERE userID = ?"
	stmt, err := u.DB.PrepareContext(ctx, SQLQUERY)
	helper.HelperIfError(err)
	_, err = stmt.ExecContext(ctx, users.Email, users.Password, users.Image, users.Nickname, users.UserID)
	helper.HelperIfError(err)
	return nil
}

func (u userRepositoryImpl) AllUser(ctx context.Context) []entity.User {
	SQLQUERY := "SELECT userID, email, nickname, password FROM yearbook_db.user"
	rows, err := u.DB.QueryContext(ctx, SQLQUERY)
	helper.HelperIfError(err)
	var users []entity.User
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.UserID, &user.Email, &user.Nickname, &user.Password)
		helper.HelperIfError(err)
		users = append(users, user)
	}
	return users
}

func (ur userRepositoryImpl) GetUserPass(ctx context.Context, email string) (uint64, string, string) {
	SQLQUERY := "SELECT userID, email, nickname, password FROM yearbook_db.user WHERE email = ?"
	rows, err := ur.DB.QueryContext(ctx, SQLQUERY, strings.ToLower(email))
	helper.HelperIfError(err)
	user := new(entity.User)
	if rows.Next() {
		err = rows.Scan(&user.UserID, &user.Email, &user.Nickname, &user.Password)
		helper.HelperIfError(err)
	} else {
		helper.NotFound("Not found", "User id "+strconv.FormatUint(user.UserID, 10)+" not found")
	}
	return user.UserID, user.Nickname, user.Password
}
