package dto

import (
	"encoding/json"
	"io"

	"github.com/deedima3/yearbook-backend/internal/blogpost/entity"
	automapper "github.com/peteprogrammer/go-automapper"
)

type BlogPostRequestBody struct {
	Content string `json:"content"`
	Pages	uint64 `json:"pages"`
}

func(b *BlogPostRequestBody) FromJSON(r io.Reader)error{
	return json.NewDecoder(r).Decode(b)
}

func(b BlogPostRequestBody) CreateBlogpostEntity() entity.Blogpost{
	var blogpostEntity entity.Blogpost
	automapper.MapLoose(b, &blogpostEntity)
	return blogpostEntity
}