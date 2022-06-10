package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type BlogPageRepository interface {
	CheckUserExist(ctx context.Context, id uint64) (bool, error)
	ViewUserPages(ctx context.Context, userID uint64) (entity.BlogPagesPeopleJoined, error)
	GetAllPages(ctx context.Context) (entity.BlogPagesPeopleJoined, error)
	CheckPages(ctx context.Context) (bool, error)
	CheckUserPage(ctx context.Context, id uint64) (bool, error)
	CreateUserPage(ctx context.Context, page entity.BlogPage) (bool, error)
	UpdateUserPage(ctx context.Context, page entity.BlogPage, pageID int) error
}
