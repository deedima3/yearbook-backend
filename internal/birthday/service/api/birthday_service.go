package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
)

type BirthdayService interface {
	GetBirthdayWeek(ctx context.Context) (dto.BlogPagesResponse, error)
	CheckUserBirthday(ctx context.Context, owner uint64) (bool, error)
}
