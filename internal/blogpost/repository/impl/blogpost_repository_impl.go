package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
)

type blogpostRepositoryImpl struct {
	DB *sql.DB
}

func ProvideBlogpostRepository(DB *sql.DB) *blogpostRepositoryImpl {
	return &blogpostRepositoryImpl{DB: DB}
}

const (
	INSERT_NEW_POST = `
	INSERT INTO railway.blogpost
	(content, pages, upvote, downvote, title)
	VALUES(?, ?, ?, ?, ?);
	`
	DELETE_POST = `
	DELETE FROM blogpost
	WHERE postID = ?;
	`
	CHECK_POST_EXISTS = `
	SELECT postID FROM blogpost
	WHERE postID = ?;
	`
	SELECT_UPVOTE_DOWNVOTE = `
	SELECT upvote, downvote FROM blogpost
	WHERE postID = ?;
	`
	SELECT_TOP_10_TWITS = `
	SELECT postID, title, content, upvote, downvote
	FROM railway.blogpost
    GROUP BY postID
    HAVING SUM(upvote) > 10
	ORDER BY upvote DESC
	LIMIT 10;
	`
	CHECK_TWITS_EXISTS = `
	SELECT postID FROM blogpost;
	`
	UPDATE_UPVOTE = `
	UPDATE blogpost SET upvote = upvote + 1
	WHERE postID = %d
	`
	UPDATE_DOWNVOTE = `
	UPDATE blogpost SET downvote = downvote + 1
	WHERE postID = ?;
	`
	GET_TWITS_PAGES = `
	SELECT postID, title, content, upvote, downvote, pages FROM blogpost
	WHERE pages = ?;
	`
	CHECK_TWITS_PAGES = `
	SELECT postID FROM blogpost
	WHERE pages = ?;
	`
)

func (b blogpostRepositoryImpl) CheckTwitsPerPages(ctx context.Context, pages uint64) (bool, error) {
	query := CHECK_TWITS_PAGES
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR CheckTwitsPerPages -> error: %v\n", err)
		return false, nil
	}
	rows, err := stmt.Query(pages)
	if err != nil {
		log.Printf("ERROR CheckTwitsPerPages -> error: %v\n", err)
		return false, nil
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (b blogpostRepositoryImpl) GetTwitsPerPages(ctx context.Context, pages uint64) (entity.BlogPosts, error) {
	query := GET_TWITS_PAGES
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR GetTwitsPerPages -> error: %v\n", err)
		return nil, err
	}
	rows, err := stmt.Query(pages)
	if err != nil {
		log.Printf("ERROR GetTwitsPerPages -> error: %v\n", err)
		return nil, err
	}

	blogposts := entity.BlogPosts{}

	for rows.Next() {
		var blogpost entity.Blogpost
		err := rows.Scan(
			&blogpost.PostId,
			&blogpost.Title,
			&blogpost.Content,
			&blogpost.Upvote,
			&blogpost.Downvote,
			&blogpost.Pages,
		)
		if err != nil {
			log.Printf("ERROR GetTwitsPerPages -> error: %v\n", err)
			return nil, err
		}
		blogposts = append(blogposts, &blogpost)
	}
	return blogposts, nil
}

func (b blogpostRepositoryImpl) CheckTwits(ctx context.Context) (bool, error) {
	query := CHECK_TWITS_EXISTS
	stmt, err := b.DB.PrepareContext(ctx, query)
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

func (b blogpostRepositoryImpl) ViewTopTwits(ctx context.Context) (entity.BlogPosts, error) {
	query := SELECT_TOP_10_TWITS
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR ViewTopTwits -> error: %v\n", err)
		return entity.BlogPosts{}, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR ViewTopTwits -> error: %v\n", err)
		return entity.BlogPosts{}, err
	}

	blogposts := entity.BlogPosts{}

	for rows.Next() {
		var blogpost entity.Blogpost

		err := rows.Scan(
			&blogpost.PostId,
			&blogpost.Title,
			&blogpost.Content,
			&blogpost.Upvote,
			&blogpost.Downvote,
		)
		if err != nil {
			log.Printf("ERROR ViewTopTwits -> error: %v\n", err)
			return entity.BlogPosts{}, err
		}
		blogposts = append(blogposts, &blogpost)
	}
	return blogposts, nil
}

func (b blogpostRepositoryImpl) UpdateUpvote(ctx context.Context, postID uint64) error {
	query := fmt.Sprintf(UPDATE_UPVOTE, postID)
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR UpdateUpvote1 -> error: %v\n", err)
		return err
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("ERROR UpdateUpvote2 -> error: %v\n", err)
		return err
	}
	return nil
}

func (b blogpostRepositoryImpl) UpdateDownvote(ctx context.Context, postID uint64) error {
	query := UPDATE_DOWNVOTE
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR UpdateDownvote -> error: %v\n", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, postID)
	if err != nil {
		log.Printf("ERROR UpdateDownvote -> error: %v\n", err)
		return err
	}
	return nil

}

func (b blogpostRepositoryImpl) ViewUpvoteDownvote(ctx context.Context, id uint64) (entity.BlogPosts, error) {
	query := SELECT_UPVOTE_DOWNVOTE
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR ViewUpvoteDownvote -> error: %v\n", err)
		return nil, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("ERROR ViewUpvoteDownvote -> error: %v\n", err)
		return nil, err
	}

	blogposts := entity.BlogPosts{}

	for rows.Next() {
		var blogpost entity.Blogpost

		err := rows.Scan(
			&blogpost.Upvote,
			&blogpost.Downvote,
		)

		if err != nil {
			log.Printf("ERROR ViewUpvoteDownvote -> error: %v\n", err)
			return nil, err
		}

		blogposts = append(blogposts, &blogpost)
	}
	return blogposts, nil
}

func (b blogpostRepositoryImpl) DeletePostByID(ctx context.Context, id uint64) error {
	query := DELETE_POST
	_, err := b.DB.Query(query, id)
	if err != nil {
		log.Printf("ERROR DeletePostByID -> error: %v\n", err)
		return err
	}
	return nil
}

func (b blogpostRepositoryImpl) CheckPostExists(ctx context.Context, id uint64) (bool, error) {
	query := CHECK_POST_EXISTS
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR CheckPostExists -> error: %v\n", err)
		return false, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("ERROR CheckPostExists -> error: %v\n", err)
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (b blogpostRepositoryImpl) InsertNewPost(ctx context.Context, bp entity.Blogpost) (uint64, error) {
	query := INSERT_NEW_POST
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR InsertNewPost -> error: %v\n", err)
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, bp.Content, bp.Pages, bp.Upvote, bp.Downvote, bp.Title)
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
