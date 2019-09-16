package mysql

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service/domain/model"
)

// UserRepository 用户资源库
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserByName(name string) (*model.User, error) {
	log.Println("[Repository@MySQL] UserRepository.GetUserByName loading...")
	var item model.User
	if err := u.db.Table("users").Where("username=?", name).Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
