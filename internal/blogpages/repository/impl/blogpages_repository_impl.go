package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
	"github.com/deedima3/yearbook-backend/internal/user/helper"
)

type blogpagesRepositoryImpl struct {
	DB *sql.DB
}

var (
	GET_ALL_PAGES = `
	SELECT bp.blogID, bp.owner, bp.header_img, u.nickname, bp.description FROM blogpages bp
	JOIN user u ON bp.owner = u.userID;
	`
	GET_USER_PAGE = `
	SELECT bp.blogID, bp.owner, bp.header_img, u.nickname, bp.description FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.userID = ?;
	`
	CHECK_USER_PAGE_EXISTS = `
	SELECT blogID FROM blogpages
	WHERE blogID = ?;
	`
	CHECK_PAGES_EXISTS = `
	SELECT blogID FROM blogpages;
	`
	CHECK_USER_EXISTS = `
	SELECT owner FROM blogpages
	WHERE owner = ?;
	`
	NEW_BLOGPAGES = `
	INSERT INTO yearbook_db.blogpages(header_img,description,owner) VALUES(?,?,?);`

	UPDATE_BLOGPAGES = `
	UPDATE yearbook_db.blogpages
	SET header_img=?,description=?
	WHERE blogID=?
	`
)

func ProvideBlogpagesRepository(DB *sql.DB) *blogpagesRepositoryImpl {
	return &blogpagesRepositoryImpl{DB: DB}
}

func (br blogpagesRepositoryImpl) CheckUserExist(ctx context.Context, id uint64) (bool, error) {
	query := CHECK_USER_EXISTS
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR CheckUserExists -> error: %v\n", err)
		return false, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("ERROR CheckUserExists -> error: %v\n", err)
		return false, err
	}

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (br blogpagesRepositoryImpl) CheckUserPage(ctx context.Context, id uint64) (bool, error) {
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
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (br blogpagesRepositoryImpl) CheckPages(ctx context.Context) (bool, error) {
	query := CHECK_PAGES_EXISTS
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR CheckPages -> error %v\n", err)
		return false, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR CheckPages -> error %v\n", err)
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (br blogpagesRepositoryImpl) GetAllPages(ctx context.Context) (entity.BlogPagesPeopleJoined, error) {
	query := GET_ALL_PAGES
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR GetAllPages -> error: %v\n", err)
		return entity.BlogPagesPeopleJoined{}, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR GetAllPages -> error: %v\n", err)
		return entity.BlogPagesPeopleJoined{}, err
	}

	blogPages := entity.BlogPagesPeopleJoined{}

	for rows.Next() {
		var blogPage entity.BlogPageUserJoined

		err := rows.Scan(
			&blogPage.BlogPage.PageID,
			&blogPage.BlogPage.Owner,
			&blogPage.BlogPage.HeaderImage,
			&blogPage.User.Nickname,
			&blogPage.BlogPage.Description,
		)

		if err != nil {
			log.Printf("ERROR GetAllPages -> error: %v\n", err)
			return entity.BlogPagesPeopleJoined{}, err
		}
		blogPages = append(blogPages, &blogPage)
	}
	return blogPages, nil
}

func (br blogpagesRepositoryImpl) ViewUserPages(ctx context.Context, userID uint64) (entity.BlogPagesPeopleJoined, error) {
	query := GET_USER_PAGE
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR ViewUserPages -> error: %v\n", err)
		return entity.BlogPagesPeopleJoined{}, err
	}
	rows, err := stmt.Query(userID)
	if err != nil {
		log.Printf("ERROR ViewUserPage -> error: %v\n", err)
		return entity.BlogPagesPeopleJoined{}, err
	}

	blogPages := entity.BlogPagesPeopleJoined{}

	for rows.Next() {
		var blogPage entity.BlogPageUserJoined

		err := rows.Scan(
			&blogPage.BlogPage.PageID,
			&blogPage.BlogPage.Owner,
			&blogPage.BlogPage.HeaderImage,
			&blogPage.User.Nickname,
			&blogPage.BlogPage.Description,
		)

		if err != nil {
			log.Printf("ERROR ViewUserPage -> error: %v\n", err)
			return entity.BlogPagesPeopleJoined{}, err
		}
		blogPages = append(blogPages, &blogPage)
	}
	return blogPages, nil
}

func (br blogpagesRepositoryImpl) CreateUserPage(ctx context.Context, page entity.BlogPage) (bool, error) {
	query := NEW_BLOGPAGES
	stmt, err := br.DB.PrepareContext(ctx, query)
	helper.HelperIfError(err)
	res, err := stmt.ExecContext(ctx, page.HeaderImage, page.Description, page.Owner)
	helper.HelperIfError(err)
	_, err = res.LastInsertId()
	helper.HelperIfError(err)
	return false, nil
}

func (br blogpagesRepositoryImpl) UpdateUserPage(ctx context.Context, page entity.BlogPage, pageID int) error {
	query := UPDATE_BLOGPAGES
	stmt, err := br.DB.PrepareContext(ctx, query)
	helper.HelperIfError(err)
	res, err := stmt.ExecContext(ctx, page.HeaderImage, page.Description, pageID)
	helper.HelperIfError(err)
	_, err = res.RowsAffected()
	helper.HelperIfError(err)
	return nil
}
