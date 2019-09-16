package service

import (
	"context"

	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/dto"
	"github.com/teamlint/mons/sample/application/query"
)

// UserService user application service
type UserService interface {
	Find(ctx context.Context, query *query.UserQuery) (*dto.User, error)
	Update(ctx context.Context, cmd *command.UpdateUserCommand) error
}
