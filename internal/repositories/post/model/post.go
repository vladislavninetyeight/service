package model

import "time"

type Post struct {
	ID        uint `bson:"post_id"`
	Detail    PostDetail
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type PostDetail struct {
	Title  string `bson:"title"`
	Body   string `bson:"body"`
	UserID uint   `bson:"user_id"`
}
