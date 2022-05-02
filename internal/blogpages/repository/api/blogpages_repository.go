package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type BlogPageRepository interface {
	ViewUserPages(ctx context.Context, userID uint64) (entity.BlogPages, error)
	CheckUserPage(ctx context.Context, id uint64)(bool, error)
}