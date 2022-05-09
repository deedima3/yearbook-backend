package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
)

type BlogpagesService interface {
	ViewUserPages(ctx context.Context, id uint64)(dto.BlogPagesResponse, error)
}