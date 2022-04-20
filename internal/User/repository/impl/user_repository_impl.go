package impl

import (
	"context"
	"database/sql"
	"github.com/deedima3/yearbook-backend/internal/User/entity"
	"github.com/deedima3/yearbook-backend/internal/User/helper"
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
	res, err := stmt.ExecContext(ctx, user.Email, user.Password, user.Nickname, user.Nim)
	helper.HelperIfError(err)
	_, err = res.LastInsertId()
	helper.HelperIfError(err)
	return nil
}
