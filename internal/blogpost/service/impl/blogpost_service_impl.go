package impl

import (
	"context"
	"log"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/blogpost/dto"
	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
	repositoryApiPkg "github.com/deedima3/yearbook-backend/internal/blogpost/repository/api"
)

type blogpostServiceImpl struct {
	rr repositoryApiPkg.BlogpostRepository
}

func ProvideRegistrationRepository(rr repositoryApiPkg.BlogpostRepository) *blogpostServiceImpl{
	return &blogpostServiceImpl{rr:rr}
}

func(bs blogpostServiceImpl)ViewUpvoteDownvote(ctx context.Context, postID uint64)(dto.UpvoteDownvoteResponses, error){
	check, err := bs.rr.CheckPostExists(ctx, postID)
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("view votes", "internal server error: "+err.Error())))
	}
	if !check || err != nil {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("post", "does not exist")))
	}
	vote, err := bs.rr.ViewUpvoteDownvote(ctx, postID)
	if err != nil {
		log.Printf("ERROR ViewUpvoteDownvote -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateUpvoteDownvoteResponses(vote), err
}

func(bs blogpostServiceImpl)DeletePostByID(ctx context.Context, postID uint64) error{
	check, err := bs.rr.CheckPostExists(ctx, postID)
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("delete post", "internal server error: "+err.Error())))
	}
	if !check || err != nil {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("post", "does not exist")))
	}
	err = bs.rr.DeletePostByID(ctx, postID)
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("delete post", "internal server error: "+err.Error())))
	}
	return nil
}

func(bs blogpostServiceImpl)CreatePost(ctx context.Context, br dto.BlogPostRequestBody)(uint64, error){
	blogID, err := bs.rr.InsertNewPost(ctx, entity.Blogpost{
		Content: br.Content,
		Pages: br.Pages,
	})
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("create post", "internal server error: "+err.Error())))
	}
	return blogID, nil	
}