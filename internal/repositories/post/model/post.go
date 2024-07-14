package model

import "time"

type Post struct {
	ID        uint
	Detail    PostDetail
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type PostDetail struct {
	Title  string
	Body   string
	UserID uint
}
