package dto

import "github.com/deedima3/yearbook-backend/internal/blogpages/entity"

type BlogPageResponse struct {
	PageID      uint64 `json:"id"`
	Owner       uint64 `json:"owner"`
	Header_img  string `json:"header_img"`
	Description string `json:"description"`
}

type BlogPagesResponse []BlogPageResponse

func CreateBlogPageResponse(bp entity.BlogPage) BlogPageResponse{
	return BlogPageResponse{
		PageID: bp.PageID,
		Owner: bp.Owner,
		Header_img: bp.HeaderImage,
		Description: bp.Description,
	}
}

func CreateBlogPagesResponse(bps entity.BlogPages) *BlogPagesResponse{
	var blogPagesResponse BlogPagesResponse
	
	for idx := range bps {
		userPage := CreateBlogPageResponse(*bps[idx])
		blogPagesResponse = append(blogPagesResponse, userPage)
	}
	return &blogPagesResponse
}