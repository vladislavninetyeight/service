package model

import "time"

type Post struct {
	UUID      string
	Info      PostDetail
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type PostDetail struct {
	FirstName string
	LastName  string
	Age       int64
}
