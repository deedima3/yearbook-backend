package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
)

type BlogpostRepository interface {
	CheckPostExists(ctx context.Context, id uint64) (bool, error)
	DeletePostByID(ctx context.Context, id uint64) error
	InsertNewPost(ctx context.Context, bp entity.Blogpost) (uint64, error)
	ViewTopTwits(ctx context.Context) (entity.BlogPosts, error)
	CheckTwits(ctx context.Context) (bool, error)
	ViewUpvoteDownvote(ctx context.Context, id uint64) (entity.BlogPosts, error)
	UpdateUpvote(ctx context.Context, postID uint64) error
	UpdateDownvote(ctx context.Context, postID uint64) error
	GetTwitsPerPages(ctx context.Context, pages uint64) (entity.BlogPosts, error)
	CheckTwitsPerPages(ctx context.Context, pages uint64) (bool, error)
}
