package mysql

import (
	"log"

	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/entity"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/repository"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

var _ = repository.UserRepository(&UserRepository{})

// UserRepository 用户资源库
type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	repo := UserRepository{}
	return &repo
}

func (u *UserRepository) GetUserByName(ctx shared.RepositoryContext, name string) (*entity.User, error) {
	log.Println("[Repository@MySQL] UserRepository.GetUserByName loading")
	var item entity.User
	if err := DB(ctx).Table("users").Where("username=?", name).Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (u *UserRepository) GetUserRoles(ctx shared.RepositoryContext, uid string) ([]*entity.Role, error) {
	log.Println("[Repository@MySQL] UserRepository.GetUserRoles loading")
	var items []*entity.Role
	if err := DB(ctx).Table("roles").Select("roles.*").
		Joins("join user_roles on user_roles.role_id=roles.id").
		Where("user_roles.user_id=?", uid).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (u *UserRepository) CreateUser(ctx shared.RepositoryContext, user *entity.User) error {
	log.Println("[Repository@MySQL] UserRepository.GetUserRoles loading")
	return DB(ctx).Create(user).Error
}
