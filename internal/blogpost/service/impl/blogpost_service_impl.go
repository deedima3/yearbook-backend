package impl

import (
	repositoryApiPkg "github.com/deedima3/yearbook-backend/internal/blogpost/repository/api"
)

type blogpostServiceImpl struct {
	rr repositoryApiPkg.BlogpostRepository
}

func ProvideRegistrationRepository(rr repositoryApiPkg.BlogpostRepository) *blogpostServiceImpl{
	return &blogpostServiceImpl{rr:rr}
}