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
	SearchUserPages(ctx context.Context, nickname string, nim string) (entity.BlogPagesPeopleJoined, error)
	GetSearchResult(ctx context.Context, nickname string, nim string) (uint64, error)
	SearchUserNickname(ctx context.Context, nickname string) (entity.BlogPagesPeopleJoined, error)
	GetSearchNickname(ctx context.Context, nickname string) (uint64, error)
	GetSearchNim(ctx context.Context, nim string) (uint64, error)
	SearchUserNim(ctx context.Context, nim string) (entity.BlogPagesPeopleJoined, error)
	UpdateUserPage(ctx context.Context, page entity.BlogPage, pageID int) error
	CheckOwnerPages(ctx context.Context, owner uint64) (uint64, error)
}
