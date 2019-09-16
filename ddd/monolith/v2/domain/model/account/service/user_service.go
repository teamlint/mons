package service

import (
	"log"

	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/entity"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/repository"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

// UserService 用户领域服务实现
type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository, ctx shared.RepositoryContext) *UserService {
	return &UserService{UserRepository: repo}
}

func (u *UserService) GetUserByName(ctx shared.RepositoryContext, name string) (*entity.User, error) {
	log.Println("[DomainService] UserService.GetUserByName loading")
	return u.UserRepository.GetUserByName(ctx, name)
}
func (u *UserService) CreateUser(ctx shared.RepositoryContext, user *entity.User) error {
	log.Println("[DomainService] UserService.CreateUser loading")
	user.Username = "DomainService-" + user.Username
	err := u.UserRepository.CreateUser(ctx, user)
	return err
}
