package dto

import (
	"encoding/json"
	"io"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
	automapper "github.com/peteprogrammer/go-automapper"
)

type BlogPostRequestBody struct {
	Content string `json:"content"`
	Pages   uint64 `json:"pages"`
}

type BlogPostVotesRequestBody struct {
	PostID uint64 `json:"postID"`
	Action string `json:"action"`
}

func (bv *BlogPostVotesRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(bv)
}

func (bv BlogPostVotesRequestBody) CreateBlogpostVotesEntity() entity.Blogpost {
	var blogpostEntity entity.Blogpost
	automapper.MapLoose(bv, &blogpostEntity)
	return blogpostEntity
}

func (b *BlogPostRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(b)
}

func (b BlogPostRequestBody) CreateBlogpostEntity() entity.Blogpost {
	var blogpostEntity entity.Blogpost
	automapper.MapLoose(b, &blogpostEntity)
	return blogpostEntity
}
