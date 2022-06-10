package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
)

type BlogpagesService interface {
	ViewUserPages(ctx context.Context, id uint64) (dto.BlogPagesResponse, error)
	GetAllPages(ctx context.Context) (dto.BlogPagesResponse, error)
	NewUserPages(ctx context.Context, blogpage dto.RequestNewBlogpage) error
	UpdateUserPages(ctx context.Context, body dto.UserUpdatePagesBody) error
}
