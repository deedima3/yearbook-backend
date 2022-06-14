package dto

import "github.com/deedima3/yearbook-backend/internal/blogpost/entity"

type UpvoteDownvoteResponse struct {
	Upvote   uint64 `json:"upvote"`
	Downvote uint64 `json:"downvote"`
}

type TopTwitsResponse struct {
	PostID   uint64 `json:"postID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Upvote   uint64 `json:"upvote"`
	Downvote uint64 `json:"downvote"`
}

type TwitsPerPagesResponse struct {
	PostID   uint64 `json:"postID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Upvote   uint64 `json:"upvote"`
	Downvote uint64 `json:"downvote"`
	Pages    uint64 `json:"pages"`
}

type TwitsPerPagesResponses []TwitsPerPagesResponse
type TopTwitsResponses []TopTwitsResponse
type UpvoteDownvoteResponses []UpvoteDownvoteResponse

func CreateTwitsPerPagesResponse(bp entity.Blogpost) TwitsPerPagesResponse {
	return TwitsPerPagesResponse{
		PostID:   bp.PostId,
		Title:    bp.Title,
		Content:  bp.Content,
		Upvote:   bp.Upvote,
		Downvote: bp.Downvote,
		Pages:    bp.Pages,
	}
}

func CreateTwitsPerPagesResponses(bp entity.BlogPosts) *TwitsPerPagesResponses {
	var twitsPerPageResponse TwitsPerPagesResponses

	for idx := range bp {
		twitsPerPages := CreateTwitsPerPagesResponse(*bp[idx])
		twitsPerPageResponse = append(twitsPerPageResponse, twitsPerPages)
	}
	return &twitsPerPageResponse
}
func CreateTopTwitsResponse(bp entity.Blogpost) TopTwitsResponse {
	return TopTwitsResponse{
		PostID:   bp.PostId,
		Title:    bp.Title,
		Content:  bp.Content,
		Upvote:   bp.Upvote,
		Downvote: bp.Downvote,
	}
}

func CreateTopTwitsResponses(bp entity.BlogPosts) *TopTwitsResponses {
	var topTwitsResponse TopTwitsResponses

	for idx := range bp {
		topTwits := CreateTopTwitsResponse(*bp[idx])
		topTwitsResponse = append(topTwitsResponse, topTwits)
	}
	return &topTwitsResponse
}

func CreateUpvoteDownvoteResponse(bp entity.Blogpost) UpvoteDownvoteResponse {
	return UpvoteDownvoteResponse{
		Upvote:   bp.Upvote,
		Downvote: bp.Downvote,
	}
}

func CreateUpvoteDownvoteResponses(bp entity.BlogPosts) *UpvoteDownvoteResponses {
	var upvoteDownvoteResponse UpvoteDownvoteResponses

	for idx := range bp {
		upvoteDownvote := CreateUpvoteDownvoteResponse(*bp[idx])
		upvoteDownvoteResponse = append(upvoteDownvoteResponse, upvoteDownvote)
	}
	return &upvoteDownvoteResponse
}
