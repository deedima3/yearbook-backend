package impl

import (
	"context"

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