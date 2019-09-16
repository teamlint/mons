package model

import (
	"time"
)

type User struct {
	ID         string
	Username   string
	Nickname   string
	Password   string
	Intro      string
	IsApproved bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type UserQuery struct {
	ID string
}
