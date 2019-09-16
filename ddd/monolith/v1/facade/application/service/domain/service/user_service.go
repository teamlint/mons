package service

import (
	"log"

	"github.com/teamlint/mons/ddd/monolith/facade/application/service/domain/model"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service/repository"
)

// UserService 用户领域服务实现
type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (u *UserService) GetUserByName(name string) (*model.User, error) {
	log.Println("[DomainService] UserService.GetUserByName loading...")
	return u.UserRepository.GetUserByName(name)
}
