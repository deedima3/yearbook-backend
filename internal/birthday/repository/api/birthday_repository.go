package api

import (
	"context"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type BirthdayRepository interface {
	CheckUserBirthday(ctx context.Context, id uint64) (bool, error)
	GetBirthdayWeek(ctx context.Context) (entity.BlogPagesPeopleJoined, error)
}
