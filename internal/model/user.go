package model

import "time"

type User struct {
	ID        uint
	detail    UserDetail
	CreatedAT time.Time
	UpdatedAT time.Time
}

type UserDetail struct {
	name  string
	phone string
}

func (u *UserDetail) GetName() string {
	return u.name
}
func (u *UserDetail) GetPhone() string {
	return u.phone
}
