package dto

import (
	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type BlogPageResponse struct {
	UserID      uint64 `json:"userID"`
	Header_img  string `json:"header_img"`
	Description string `json:"description"`
	Nickname    string `json:"nickname"`
	Nim         string `json:"nim"`
	Image       string `json:"image"`
}

type BlogPagesSearchResponse struct {
	UserID      uint64 `json:"userID"`
	Header_img  string `json:"header_img"`
	Description string `json:"description"`
	Nickname    string `json:"nickname"`
	Nim         string `json:"nim"`
	Image       string `json:"image"`
}

type BlogPagesSearchResponses []BlogPagesSearchResponse
type BlogPagesResponse []BlogPageResponse

func CreateBlogPagesSearchResponse(bp entity.BlogPageUserJoined) BlogPagesSearchResponse {
	return BlogPagesSearchResponse{
		UserID:      bp.User.UserID,
		Header_img:  bp.BlogPage.HeaderImage,
		Description: bp.BlogPage.Description,
		Nickname:    bp.User.Nickname,
		Nim:         bp.User.Nim,
		Image:       bp.User.Image,
	}
}

func CreateBlogPagesSearchResponses(bps entity.BlogPagesPeopleJoined) *BlogPagesSearchResponses {
	var blogPageSearchResponse BlogPagesSearchResponses

	for idx := range bps {
		searchRes := CreateBlogPagesSearchResponse(*bps[idx])
		blogPageSearchResponse = append(blogPageSearchResponse, searchRes)
	}
	return &blogPageSearchResponse
}

func CreateBlogPageResponse(bp entity.BlogPageUserJoined) BlogPageResponse {
	return BlogPageResponse{
		UserID:      bp.User.UserID,
		Header_img:  bp.BlogPage.HeaderImage,
		Description: bp.BlogPage.Description,
		Nickname:    bp.User.Nickname,
		Nim:         bp.User.Nim,
		Image:       bp.User.Image,
	}
}

func CreateBlogPagesResponse(bps entity.BlogPagesPeopleJoined) *BlogPagesResponse {
	var blogPagesResponse BlogPagesResponse

	for idx := range bps {
		userPage := CreateBlogPageResponse(*bps[idx])
		blogPagesResponse = append(blogPagesResponse, userPage)
	}
	return &blogPagesResponse
}
