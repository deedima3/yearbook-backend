package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
	"github.com/deedima3/yearbook-backend/internal/user/helper"
)

type blogpagesRepositoryImpl struct {
	DB *sql.DB
}

var (
	GET_ALL_PAGES = `
	SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
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
	INSERT INTO railway.blogpages(header_img,description,owner) VALUES(?,?,?);
	`
	SEARCH_BLOGPAGE = `
	SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%'
	UNION
	SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nim LIKE '%s%%';
	`
	COUNT_SEARCH_RES = `
	SELECT COUNT(u.userID) FROM blogpages AS bp
	JOIN user as u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%' OR u.nim LIKE '%s%%';
	`
	SEARCH_NICKNAME = `
	SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%';
	`
	COUNT_SEARCH_NICKNAME = `
	SELECT COUNT(u.userID) FROM blogpages AS bp
	JOIN user as u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%';
	`
	SEARCH_NIM = `
	SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nim LIKE '%s%%';
	`
	COUNT_SEARCH_NIM = `
	SELECT COUNT(u.userID) FROM blogpages AS bp
	JOIN user as u ON bp.owner = u.userID
	WHERE u.nim LIKE '%s%%';
  `
	UPDATE_BLOGPAGES = `
	UPDATE railway.blogpages
	SET header_img=?,description=?
	WHERE blogID=?
	`
	CHECK_OWNER_PAGES = `
	SELECT COUNT(blogID) FROM blogpages
	WHERE owner IN (?);
	`
)

func ProvideBlogpagesRepository(DB *sql.DB) *blogpagesRepositoryImpl {
	return &blogpagesRepositoryImpl{DB: DB}
}

func (br blogpagesRepositoryImpl) CheckOwnerPages(ctx context.Context, owner uint64) (uint64, error) {
	query := CHECK_OWNER_PAGES
	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR CheckOwnerPages -> error: %v\n", err)
		return 0, err
	}
	rows, err := stmt.Query(owner)
	if err != nil {
		log.Printf("ERROR CheckOwnerPages -> error: %v\n", err)
		return 0, err
	}
	var ownerCount uint64
	for rows.Next() {
		err := rows.Scan(&ownerCount)
		if err != nil {
			log.Printf("ERROR CheckOwnerPages -> error: %v\n", err)
			return 0, err
		}
	}
	return ownerCount, nil
}

func (br blogpagesRepositoryImpl) SearchUserNim(ctx context.Context, nim string) (entity.BlogPagesPeopleJoined, error) {
	query := SEARCH_NIM
	querySearch := fmt.Sprintf(query, nim)

	stmt, err := br.DB.PrepareContext(ctx, querySearch)
	if err != nil {
		log.Printf("ERROR SearchUserNim -> error: %v\n", err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR SearchUserNim -> error: %v\n", err)
		return nil, err
	}

	blogPagesJoined := entity.BlogPagesPeopleJoined{}

	for rows.Next() {
		var blogPageJoined entity.BlogPageUserJoined

		err := rows.Scan(
			&blogPageJoined.User.UserID,
			&blogPageJoined.BlogPage.HeaderImage,
			&blogPageJoined.BlogPage.Description,
			&blogPageJoined.User.Nickname,
			&blogPageJoined.User.Nim,
			&blogPageJoined.User.Image,
		)

		if err != nil {
			log.Printf("ERROR SearchUserNim -> error: %v\n", err)
			return nil, err
		}
		blogPagesJoined = append(blogPagesJoined, &blogPageJoined)
	}
	return blogPagesJoined, nil
}

func (br blogpagesRepositoryImpl) GetSearchNim(ctx context.Context, nim string) (uint64, error) {
	query := COUNT_SEARCH_NIM
	queryRes := fmt.Sprintf(query, nim)

	stmt, err := br.DB.PrepareContext(ctx, queryRes)
	if err != nil {
		log.Printf("ERROR GetSearchNim -> error: %v\n", err)
		return 0, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR GetSearchNim -> error: %v\n", err)
		return 0, err
	}

	var searchCount uint64
	for rows.Next() {
		err := rows.Scan(
			&searchCount,
		)
		if err != nil {
			log.Printf("ERROR GetSearchNim -> error: %v\n", err)
			return 0, nil
		}
	}
	return searchCount, nil
}

func (br blogpagesRepositoryImpl) SearchUserNickname(ctx context.Context, nickname string) (entity.BlogPagesPeopleJoined, error) {
	query := SEARCH_NICKNAME
	querySearch := fmt.Sprintf(query, nickname)

	stmt, err := br.DB.PrepareContext(ctx, querySearch)
	if err != nil {
		log.Printf("ERROR SearchUserNickname -> error: %v\n", err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR SearchUserNickname -> error: %v\n", err)
		return nil, err
	}

	blogPagesJoined := entity.BlogPagesPeopleJoined{}

	for rows.Next() {
		var blogPageJoined entity.BlogPageUserJoined

		err := rows.Scan(
			&blogPageJoined.User.UserID,
			&blogPageJoined.BlogPage.HeaderImage,
			&blogPageJoined.BlogPage.Description,
			&blogPageJoined.User.Nickname,
			&blogPageJoined.User.Nim,
			&blogPageJoined.User.Image,
		)

		if err != nil {
			log.Printf("ERROR SearchUserNickname -> error: %v\n", err)
			return nil, err
		}
		blogPagesJoined = append(blogPagesJoined, &blogPageJoined)
	}
	return blogPagesJoined, nil
}

func (br blogpagesRepositoryImpl) GetSearchNickname(ctx context.Context, nickname string) (uint64, error) {
	query := COUNT_SEARCH_RES
	queryRes := fmt.Sprintf(query, nickname)

	stmt, err := br.DB.PrepareContext(ctx, queryRes)
	if err != nil {
		log.Printf("ERROR GetSearchNickname -> error: %v\n", err)
		return 0, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR GetSearchNickname -> error: %v\n", err)
		return 0, err
	}

	var searchCount uint64
	for rows.Next() {
		err := rows.Scan(
			&searchCount,
		)
		if err != nil {
			log.Printf("ERROR GetSearchNickname -> error: %v\n", err)
			return 0, nil
		}
	}
	return searchCount, nil
}

func (br blogpagesRepositoryImpl) GetSearchResult(ctx context.Context, nickname string, nim string) (uint64, error) {
	query := COUNT_SEARCH_RES
	queryRes := fmt.Sprintf(query, nickname, nim)

	stmt, err := br.DB.PrepareContext(ctx, queryRes)
	if err != nil {
		log.Printf("ERROR GetSearchResult -> error: %v\n", err)
		return 0, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR GetSearchResult -> error: %v\n", err)
		return 0, err
	}

	var searchCount uint64
	for rows.Next() {
		err := rows.Scan(
			&searchCount,
		)
		if err != nil {
			log.Printf("ERROR GetSearchResult -> error: %v\n", err)
			return 0, nil
		}
	}
	return searchCount, nil
}

func (br blogpagesRepositoryImpl) SearchUserPages(ctx context.Context, nickname string, nim string) (entity.BlogPagesPeopleJoined, error) {
	query := SEARCH_BLOGPAGE
	querySearch := fmt.Sprintf(query, nickname, nim)

	stmt, err := br.DB.PrepareContext(ctx, querySearch)
	if err != nil {
		log.Printf("ERROR SearchUserPages -> error: %v\n", err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR SearchUserPages -> error: %v\n", err)
		return nil, err
	}

	blogPagesJoined := entity.BlogPagesPeopleJoined{}

	for rows.Next() {
		var blogPageJoined entity.BlogPageUserJoined

		err := rows.Scan(
			&blogPageJoined.User.UserID,
			&blogPageJoined.BlogPage.HeaderImage,
			&blogPageJoined.BlogPage.Description,
			&blogPageJoined.User.Nickname,
			&blogPageJoined.User.Nim,
			&blogPageJoined.User.Image,
		)

		if err != nil {
			log.Printf("ERROR SearchUserPages -> error: %v\n", err)
			return nil, err
		}
		blogPagesJoined = append(blogPagesJoined, &blogPageJoined)
	}
	return blogPagesJoined, nil
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
			&blogPage.User.UserID,
			&blogPage.BlogPage.HeaderImage,
			&blogPage.BlogPage.Description,
			&blogPage.User.Nickname,
			&blogPage.User.Nim,
			&blogPage.User.Image,
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
