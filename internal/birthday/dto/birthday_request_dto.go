package dto

import (
	"encoding/json"
	"io"
)

type RequestCheckUserBirthday struct {
	Owner int64 `json:"owner"`
}

func (u *RequestCheckUserBirthday) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}
