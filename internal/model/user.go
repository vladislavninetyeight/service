package model

import "time"

type User struct {
	ID        uint
	Detail    UserDetail
	CreatedAT time.Time
	UpdatedAT time.Time
}

type UserDetail struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
