package impl

import (
	"database/sql"
)

type blogpostRepositoryImpl struct {
	DB *sql.DB
}

func ProvideBlogpostRepository(DB *sql.DB) *blogpostRepositoryImpl {
	return &blogpostRepositoryImpl{DB:DB}
}