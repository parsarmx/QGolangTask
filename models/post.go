package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PostRequest struct {
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (pr PostRequest) Validate() error {

	return validation.ValidateStruct(
		&pr,
		validation.Field(&pr.Title, validation.Required),
		validation.Field(&pr.Content, validation.Required),
	)
}