package user

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Password   string    `json:"password"`
	Intro      string    `json:"intro"`
	IsApprvoed bool      `json:"is_apprvoed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
