package service

import (
	"context"
    "fmt"

    api "github.com/teamlint/mons/sample2/services/user/api"
)

type userService struct{}

func NewUserService() api.UserServer {
    return &userService{}
}

func New(mdw []Middleware) api.UserServer {
    svc := NewUserService()
	for _, m := range mdw {
		svc = m(svc)
	}
	return svc
}

func (s *userService) Find(ctx context.Context, req *api.FindUserRequest) (*api.FindUserReply, error) {
    // TODO
	return nil, fmt.Errorf("not implemented")
}

func (s *userService) Update(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserReply, error) {
    // TODO
	return nil, fmt.Errorf("not implemented")
}

