package entity

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Password   string    `json:"-"`
	Intro      string    `json:"intro"`
	IsApproved bool      `json:"is_approved"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
