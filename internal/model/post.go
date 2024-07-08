package model

import "time"

type Post struct {
	ID        uint
	Detail    PostDetail
	CreatedAT time.Time
	UpdatedAT time.Time
}

type PostDetail struct {
	title  string
	body   string
	userID uint
}

func (p *PostDetail) GetTitle() string { return p.title }
func (p *PostDetail) GetBody() string  { return p.body }
func (p *PostDetail) GetUserID() uint  { return p.userID }
