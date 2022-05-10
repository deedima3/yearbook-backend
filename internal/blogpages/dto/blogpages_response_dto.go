package dto

import "github.com/deedima3/yearbook-backend/internal/blogpages/entity"

type BlogPageResponse struct {
	PageID      uint64 `json:"id"`
	Owner       uint64 `json:"owner"`
	Nickname	string `json:"nickname"`
	Header_img  string `json:"header_img"`
	Description string `json:"description"`
}

type BlogPagesResponse []BlogPageResponse

func CreateBlogPageResponse(bp entity.BlogPageUserJoined) BlogPageResponse{
	return BlogPageResponse{
		PageID: bp.BlogPage.PageID,
		Owner: bp.BlogPage.Owner,
		Nickname: bp.User.Nickname,
		Header_img: bp.BlogPage.HeaderImage,
		Description: bp.BlogPage.Description,
	}
}

func CreateBlogPagesResponse(bps entity.BlogPagesPeopleJoined) *BlogPagesResponse{
	var blogPagesResponse BlogPagesResponse
	
	for idx := range bps {
		userPage := CreateBlogPageResponse(*bps[idx])
		blogPagesResponse = append(blogPagesResponse, userPage)
	}
	return &blogPagesResponse
}