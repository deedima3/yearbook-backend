package impl

import (
	"context"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
	"github.com/deedima3/yearbook-backend/internal/user/helper"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/blogpages/dto"
	repositoryApiPkg "github.com/deedima3/yearbook-backend/internal/blogpages/repository/api"
)

type blogpageServiceImpl struct {
	rr repositoryApiPkg.BlogPageRepository
}

func ProvideRegistrationRepository(rr repositoryApiPkg.BlogPageRepository) *blogpageServiceImpl {
	return &blogpageServiceImpl{rr: rr}
}

func (bp blogpageServiceImpl) SearchUserNim(ctx context.Context, nim string) (dto.BlogPagesSearchResponses, error) {
	searchCount, err := bp.rr.GetSearchNim(ctx, nim)
	if err != nil {
		log.Printf("ERROR SearchUserNim -> error: %v\n", err)
		return nil, err
	}
	if searchCount == 0 {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("search", "not found")))
	}
	searchRes, err := bp.rr.SearchUserNim(ctx, nim)
	if err != nil {
		log.Printf("ERROR SearchUserNim -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateBlogPagesSearchResponses(searchRes), nil
}

func (bp blogpageServiceImpl) SearchUserNickname(ctx context.Context, nickname string) (dto.BlogPagesSearchResponses, error) {
	searchCount, err := bp.rr.GetSearchNickname(ctx, nickname)
	if err != nil {
		log.Printf("ERROR SearchUserNickname -> error: %v\n", err)
		return nil, err
	}
	if searchCount == 0 {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("search", "not found")))
	}
	searchRes, err := bp.rr.SearchUserNickname(ctx, nickname)
	if err != nil {
		log.Printf("ERROR SearchUserNickname -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateBlogPagesSearchResponses(searchRes), nil
}

func (bp blogpageServiceImpl) SearchUserPages(ctx context.Context, nickname string, nim string) (dto.BlogPagesSearchResponses, error) {
	searchCount, err := bp.rr.GetSearchResult(ctx, nickname, nim)
	if err != nil {
		log.Printf("ERROR SearchUserPages -> error: %v\n", err)
		return nil, err
	}
	if searchCount == 0 {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("search", "not found")))
	}
	searchRes, err := bp.rr.SearchUserPages(ctx, nickname, nim)
	if err != nil {
		log.Printf("ERROR SearchUserPages -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateBlogPagesSearchResponses(searchRes), nil
}

func (bp blogpageServiceImpl) GetAllPages(ctx context.Context) (dto.BlogPagesResponse, error) {
	check, err := bp.rr.CheckPages(ctx)
	if err != nil {
		log.Printf("ERROR GetAllPages -> error: %v\n", err)
		return nil, err
	}
	if !check {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("pages", "does not exist")))
	}
	allPages, err := bp.rr.GetAllPages(ctx)
	if err != nil {
		log.Printf("ERROR GetAllPages -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateBlogPagesResponse(allPages), err
}

func (bp blogpageServiceImpl) ViewUserPages(ctx context.Context, id uint64) (dto.BlogPagesResponse, error) {
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

func (bp blogpageServiceImpl) NewUserPages(ctx context.Context, blogpage dto.RequestNewBlogpage) error {
	page := entity.BlogPage{
		HeaderImage: blogpage.Header_img_path,
		Description: blogpage.Description,
		Owner:       uint64(blogpage.UserID),
	}
	sum, err := bp.rr.CheckOwnerPages(ctx, page.Owner)
	helper.HelperIfError(err)
	if sum == 0 {
		_, err = bp.rr.CreateUserPage(ctx, page)
		helper.HelperIfError(err)
	}
	if sum > 0 {
		panic(sicgolib.NewErrorResponse(503, "service unavailable",
			sicgolib.NewErrorResponseValue("post pages service", "unavailable")))
	}
	return nil
}

func (bp blogpageServiceImpl) UpdateUserPages(ctx context.Context, body dto.UserUpdatePagesBody) error {
	pagesUpdate := entity.BlogPage{
		HeaderImage: body.HeaderImgPath,
		Description: body.Description,
	}
	err := bp.rr.UpdateUserPage(ctx, pagesUpdate, body.BlogID)
	helper.HelperIfError(err)
	return nil
}
