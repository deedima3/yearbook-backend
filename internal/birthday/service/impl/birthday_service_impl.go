package impl

import (
	"context"
	"log"

	repositoryApiPkg "github.com/deedima3/yearbook-backend/internal/birthday/repository/api"
	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
)

type birthdayServiceImpl struct {
	rr repositoryApiPkg.BirthdayRepository
}

func ProvideRegistrationRepository(rr repositoryApiPkg.BirthdayRepository) *birthdayServiceImpl {
	return &birthdayServiceImpl{rr: rr}
}

func (bp birthdayServiceImpl) GetBirthdayWeek(ctx context.Context) (dto.BlogPagesResponse, error) {
	birthdayWeek, err := bp.rr.GetBirthdayWeek(ctx)
	if err != nil {
		log.Printf("ERROR GetBirthdayWeek -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateBlogPagesResponse(birthdayWeek), err
}

func (bp birthdayServiceImpl) CheckUserBirthday(ctx context.Context, owner uint64) (bool, error) {
	return bp.rr.CheckUserBirthday(ctx, owner)
}
