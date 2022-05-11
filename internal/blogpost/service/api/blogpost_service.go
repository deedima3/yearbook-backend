package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpost/dto"
)

type BlogpostService interface {
	DeletePostByID(ctx context.Context, postID uint64) error
	CreatePost(ctx context.Context, br dto.BlogPostRequestBody) (uint64, error)
}