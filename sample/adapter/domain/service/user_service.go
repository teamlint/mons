package service

import (
	dosvc "github.com/teamlint/mons/sample/domain/service"
)

type UserService struct{}

func NewUserService() dosvc.UserService {
	return &UserService{}
}
func (s *UserService) Duplicated(username string) error {
	return nil
}
