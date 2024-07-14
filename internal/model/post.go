package model

import "time"

type Post struct {
	ID        uint
	Detail    PostDetail
	CreatedAT time.Time
	UpdatedAT time.Time
}

type PostDetail struct {
	Title  string
	Body   string
	UserID uint
}
