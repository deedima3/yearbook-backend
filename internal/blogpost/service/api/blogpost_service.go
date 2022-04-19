package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpost/dto"
)

type BlogpostService interface {
	CreatePost(ctx context.Context, br dto.BlogPostRequestBody) (uint64, error)
}