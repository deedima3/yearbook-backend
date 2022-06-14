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

func ProvideRegistrationRepository(rr repositoryApiPkg.BlogpostRepository) *blogpostServiceImpl {
	return &blogpostServiceImpl{rr: rr}
}

func (bs blogpostServiceImpl) GetTwitsPerPages(ctx context.Context, pages uint64) (dto.TwitsPerPagesResponses, error) {
	check, err := bs.rr.CheckTwitsPerPages(ctx, pages)
	if err != nil {
		log.Printf("ERROR GetTwitsPerPages -> error: %v\n", err)
		return nil, err
	}
	if !check || err != nil {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("twitss", "does not exist")))
	}
	twits, err := bs.rr.GetTwitsPerPages(ctx, pages)
	if err != nil {
		log.Printf("ERROR GetTwitsPerPages -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateTwitsPerPagesResponses(twits), nil
}

func (bs blogpostServiceImpl) ViewUpvoteDownvote(ctx context.Context, postID uint64) (dto.UpvoteDownvoteResponses, error) {
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
	return *dto.CreateUpvoteDownvoteResponses(vote), nil
}

func (bs blogpostServiceImpl) ViewTopTwits(ctx context.Context) (dto.TopTwitsResponses, error) {
	check, err := bs.rr.CheckTwits(ctx)
	if err != nil {
		log.Printf("ERROR ViewTopTwits -> error: %v\n", err)
		return nil, err
	}
	if !check || err != nil {
		panic(sicgolib.NewErrorResponse(404, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("twits", "does not exist")))
	}
	topTwits, err := bs.rr.ViewTopTwits(ctx)
	if err != nil {
		log.Printf("ERROR ViewTopTwits -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateTopTwitsResponses(topTwits), nil
}

func (bs blogpostServiceImpl) DeletePostByID(ctx context.Context, postID uint64) error {
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

func (bs blogpostServiceImpl) CreatePost(ctx context.Context, br dto.BlogPostRequestBody) (uint64, error) {
	blogID, err := bs.rr.InsertNewPost(ctx, entity.Blogpost{
		Content:  br.Content,
		Pages:    br.Pages,
		Upvote:   br.Upvote,
		Downvote: br.Downvote,
		Title:    br.Title,
	})
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("create post", "internal server error: "+err.Error())))
	}
	return blogID, nil
}

func (bs blogpostServiceImpl) UpdateVotes(ctx context.Context, bv dto.BlogPostVotesRequestBody) (string, error) {
	if bv.Action == "up" {
		err := bs.rr.UpdateUpvote(ctx, bv.PostID)
		if err != nil {
			panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
				sicgolib.NewErrorResponseValue("upvote", "internal server error: "+err.Error())))
		}
		return "upvote sukses", nil
	} else if bv.Action == "down" {
		err := bs.rr.UpdateDownvote(ctx, bv.PostID)
		if err != nil {
			panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
				sicgolib.NewErrorResponseValue("downvote", "internal server error: "+err.Error())))
		}
		return "downvote sukses", nil
	}
	return "update sukses", nil
}
