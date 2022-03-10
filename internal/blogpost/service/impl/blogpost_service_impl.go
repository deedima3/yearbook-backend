package impl

import (
	"context"

	repositoryApiPkg "github.com/deedima3/yearbook-backend/internal/blogpost/repository/api"
)

type blogpostServiceImpl struct {
	rr repositoryApiPkg.BlogpostRepository
}

func VoteTwits(ctx context.Context, id uint64)()

func ProvideRegistrationRepository(rr repositoryApiPkg.BlogpostRepository) *blogpostServiceImpl{
	return &blogpostServiceImpl{rr:rr}
}