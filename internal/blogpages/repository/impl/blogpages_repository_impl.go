package impl

import "database/sql"

type blogpagesRepositoryImpl struct {
	DB *sql.DB
}

func ProvideBlogpagesRepository(DB *sql.DB) *blogpagesRepositoryImpl{
	return &blogpagesRepositoryImpl{DB: DB}
}