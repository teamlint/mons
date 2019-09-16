package mysql

import (
	"log"

	"github.com/teamlint/mons/sample/domain/model"
	"github.com/teamlint/mons/sample/domain/repository"
	"github.com/teamlint/mons/shared/adapter/repository/mysql"
	shared "github.com/teamlint/mons/shared/domain/repository"
)

var _ = repository.UserRepository(&UserRepository{})

// UserRepository 用户资源库
type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	repo := UserRepository{}
	return &repo
}

func (u *UserRepository) Create(ctx shared.RepositoryContext, user *model.User) error {
	log.Println("[Repository@MySQL] UserRepository.Create loading")
	return mysql.DB(ctx).Table("users").Create(user).Error
}
func (u *UserRepository) FindByID(ctx shared.RepositoryContext, id string) (*model.User, error) {
	log.Println("[Repository@MySQL] UserRepository.FindByID loading")
	var item model.User
	if err := mysql.DB(ctx).Table("users").Where("id=?", id).Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (u *UserRepository) FindAll(ctx shared.RepositoryContext, pageIndex int, pageSize int) ([]*model.User, int, error) {
	log.Println("[Repository@MySQL] UserRepository.FindAll loading")
	var items []*model.User
	var total int
	offset := pageIndex*pageSize - pageSize
	err := mysql.DB(ctx).Table("users").Count(&total).Offset(offset).Limit(pageSize).Scan(&items).Error
	return items, total, err
}
func (u *UserRepository) Update(ctx shared.RepositoryContext, user *model.User) error {
	log.Println("[Repository@MySQL] UserRepository.Update loading")
	return mysql.DB(ctx).Table("users").Omit("created_at", "updated_at").Save(user).Error
}
func (u *UserRepository) Delete(ctx shared.RepositoryContext, id string) error {
	log.Println("[Repository@MySQL] UserRepository.Delete loading")
	return mysql.DB(ctx).Table("users").Where("id=?", id).Delete(&model.User{}).Error
}
