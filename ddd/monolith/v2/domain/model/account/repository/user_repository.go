package repository

import (
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/entity"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

// UserRepository 用户资源库接口
type UserRepository interface {
	GetUserByName(ctx shared.RepositoryContext, name string) (*entity.User, error)
	GetUserRoles(ctx shared.RepositoryContext, id string) ([]*entity.Role, error)
	CreateUser(ctx shared.RepositoryContext, user *entity.User) error
}
