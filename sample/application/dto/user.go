package dto

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/teamlint/mons/sample/domain/model"
)

// User application user info
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

func (u *User) From(do *model.User) error {
	return copier.Copy(&u, do)
}
