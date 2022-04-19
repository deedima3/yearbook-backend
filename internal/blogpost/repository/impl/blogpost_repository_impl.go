package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
)

type blogpostRepositoryImpl struct {
	DB *sql.DB
}

func ProvideBlogpostRepository(DB *sql.DB) *blogpostRepositoryImpl {
	return &blogpostRepositoryImpl{DB:DB}
}

const(
	INSERT_NEW_POST = `
	INSERT INTO yearbook_db.blogpost
	(content, pages)
	VALUES(?, ?);
	`
)

func(b blogpostRepositoryImpl)InsertNewPost(ctx context.Context, bp entity.Blogpost)(uint64, error){
	query := INSERT_NEW_POST
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR InsertNewPost -> error: %v\n", err)
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, bp.Content, bp.Pages)
	if err != nil {
		log.Printf("ERROR InsertNewPost -> error: %v\n", err)
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("ERROR InsertNewPost -> error: %v\n", err)
		return 0, err
	}
	return uint64(lastID), nil
}