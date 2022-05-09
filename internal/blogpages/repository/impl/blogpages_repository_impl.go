package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type blogpagesRepositoryImpl struct {
	DB *sql.DB
}

var(
	GET_USER_PAGE = `
	SELECT bp.blogID, bp.owner,bp.header_img, bp.description FROM blogpages bp
	WHERE bp.blogID = ?
	`
	CHECK_USER_PAGE_EXISTS = `
	SELECT blogID FROM blogpages
	WHERE blogID = ?
	`
)

func ProvideBlogpagesRepository(DB *sql.DB) *blogpagesRepositoryImpl{
	return &blogpagesRepositoryImpl{DB: DB}
}

func(br blogpagesRepositoryImpl)CheckUserPage(ctx context.Context, id uint64)(bool, error){
	query := CHECK_USER_PAGE_EXISTS
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR CheckUserPage -> error %v\n", err)
		return false, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("ERROR CheckUserPage -> error %v\n", err)
		return false, err
	}
	if rows.Next(){
		return true, nil
	}
	return false, nil
}

func(br blogpagesRepositoryImpl)ViewUserPages(ctx context.Context, userID uint64)(entity.BlogPages, error){
	query := GET_USER_PAGE
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR ViewUserPages -> error: %v\n", err)
		return entity.BlogPages{}, err
	}
	rows, err := stmt.Query(userID)
	if err != nil {
		log.Printf("ERROR ViewUserPage -> error: %v\n", err)
		return entity.BlogPages{}, err
	}

	blogPages := entity.BlogPages{}

	for rows.Next(){
		var blogPage entity.BlogPage
		
		err := rows.Scan(&blogPage.PageID, &blogPage.Owner, &blogPage.HeaderImage, &blogPage.Description)
		if err != nil {
			log.Printf("ERROR ViewUserPage -> error: %v\n", err)
			return entity.BlogPages{}, err
		}
		blogPages = append(blogPages, &blogPage)
	}
	return blogPages, nil
}
