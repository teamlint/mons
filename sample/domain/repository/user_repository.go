package repository

import (
	"github.com/teamlint/mons/sample/domain/model"
	shared "github.com/teamlint/mons/shared/domain/repository"
)

type UserRepository interface {
	Create(rc shared.RepositoryContext, user *model.User) error
	FindByID(rc shared.RepositoryContext, id string) (*model.User, error)
	FindAll(rc shared.RepositoryContext, pageIndex int, pageSize int) ([]*model.User, int, error)
	Update(rc shared.RepositoryContext, user *model.User) error
	Delete(rc shared.RepositoryContext, id string) error
}
