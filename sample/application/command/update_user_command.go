package command

import (
	"github.com/jinzhu/copier"
	"github.com/teamlint/mons/sample/application/dto"
)

type UpdateUserCommand struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Intro    string `json:"intro"`
}

func NewUpdateUserCommand(id, username, nickname, pwd, intro string) *UpdateUserCommand {
	return &UpdateUserCommand{
		ID:       id,
		Username: username,
		Nickname: nickname,
		Password: pwd,
		Intro:    intro,
	}
}
func UpdateUserCommandFrom(user *dto.User) *UpdateUserCommand {
	var cmd UpdateUserCommand
	if user != nil {
		if err := copier.Copy(&cmd, user); err == nil {
			return &cmd
		}
	}
	return nil
}
