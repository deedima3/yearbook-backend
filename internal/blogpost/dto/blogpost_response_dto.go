package dto

import "github.com/deedima3/yearbook-backend/internal/blogpost/entity"

type UpvoteDownvoteResponse struct {
	Upvote   uint64 `json:"upvote"`
	Downvote uint64 `json:"downvote"`
}

type UpvoteDownvoteResponses []UpvoteDownvoteResponse

func CreateUpvoteDownvoteResponse(bp entity.Blogpost)UpvoteDownvoteResponse{
	return UpvoteDownvoteResponse{
		Upvote: bp.Upvote,
		Downvote: bp.Downvote,
	}
}

func CreateUpvoteDownvoteResponses(bp entity.BlogPosts) *UpvoteDownvoteResponses{
	var upvoteDownvoteResponse UpvoteDownvoteResponses

	for idx := range bp {
		upvoteDownvote := CreateUpvoteDownvoteResponse(*bp[idx])
		upvoteDownvoteResponse = append(upvoteDownvoteResponse, upvoteDownvote)
	}
	return &upvoteDownvoteResponse
}