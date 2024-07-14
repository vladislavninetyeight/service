package model

import "time"

type PostGetAllResponse struct {
	Posts []Post
}

type Post struct {
	ID        uint
	Detail    PostDetail
	CreatedAT time.Time
	UpdatedAT time.Time
}

type PostDetail struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"`
}
