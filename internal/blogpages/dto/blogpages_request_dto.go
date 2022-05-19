package dto

import (
	"encoding/json"
	"io"
)

type RequestNewBlogpage struct {
	UserID          int64  `json:"userID"`
	Header_img_path string `json:"header_Img_Path"`
	Description     string `json:"description"`
}

func (u *RequestNewBlogpage) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}
