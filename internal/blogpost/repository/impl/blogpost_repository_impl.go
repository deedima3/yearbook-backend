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
	INSERT INTO yearbook_db.blogpost
	(content, pages)
	VALUES(?, ?);
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
	UPDATE_UPVOTE = `
	UPDATE blogpost SET upvote = upvote + 1
	WHERE postID = %d
	`
	UPDATE_DOWNVOTE = `
	UPDATE blogpost SET downvote = downvote + 1
	WHERE postID = ?;
	`
)

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
