package impl

import (
	"context"
	"log"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
	repositoryApiPkg "github.com/deedima3/yearbook-backend/internal/blogpages/repository/api"
)

type blogpageServiceImpl struct {
	rr repositoryApiPkg.BlogPageRepository
}

func ProvideRegistrationRepository(rr repositoryApiPkg.BlogPageRepository) *blogpageServiceImpl{	
	return &blogpageServiceImpl{rr: rr}
}

func(bp blogpageServiceImpl)ViewUserPages(ctx context.Context, id uint64)(dto.BlogPagesResponse, error){
	checkUser, err := bp.rr.CheckUserExist(ctx, id)
	if err != nil {
		log.Printf("ERROR ViewUserPages -> error: %v\n", err)
		return nil, err
	}

	if !checkUser || err != nil {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("user", "does not exist")))
	}

	checkPages, err := bp.rr.CheckUserPage(ctx, id)
	if err != nil {
		log.Printf("ERROR ViewUserpages -> error: %v\n", err)
		return nil, err
	}

	if !checkPages || err != nil {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("user pages", "does not exist")))
	}

	userPages, err := bp.rr.ViewUserPages(ctx, id)
	if err != nil {
		log.Printf("ERROR ViewUserPage -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateBlogPagesResponse(userPages), err
}