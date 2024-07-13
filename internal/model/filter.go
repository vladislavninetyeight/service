package model

import "time"

type Filter struct {
	FromCreatedAt  *time.Time `json:"fromCreatedAt"`
	ToCreatedAt    *time.Time `json:"toCreatedAt"`
	Name           string     `json:"name"`
	Limit          uint       `json:"limit"`
	Offset         uint       `json:"offset"`
	TopPostsAmount string     `json:"topPostsAmount"`
}
