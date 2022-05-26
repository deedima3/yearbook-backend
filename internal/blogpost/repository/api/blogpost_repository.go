package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
)

type BlogpostRepository interface {
	CheckPostExists(ctx context.Context, id uint64)(bool, error)
	DeletePostByID(ctx context.Context, id uint64)error
	InsertNewPost(ctx context.Context, bp entity.Blogpost)(uint64, error)
	ViewUpvoteDownvote(ctx context.Context, id uint64) (entity.BlogPosts, error)
}