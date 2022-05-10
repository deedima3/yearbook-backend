package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type BlogPageRepository interface {
	CheckUserExist(ctx context.Context, id uint64)(bool, error)
	ViewUserPages(ctx context.Context, userID uint64)(entity.BlogPagesPeopleJoined, error)
	CheckUserPage(ctx context.Context, id uint64)(bool, error)
}