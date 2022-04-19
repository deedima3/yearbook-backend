package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
)

type BlogpostRepository interface {
	InsertNewPost(ctx context.Context, bp entity.Blogpost)(uint64, error)
}