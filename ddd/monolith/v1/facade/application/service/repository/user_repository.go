package repository

import "github.com/teamlint/mons/ddd/monolith/facade/application/service/domain/model"

// UserRepository 用户资源库接口
type UserRepository interface {
	GetUserByName(name string) (*model.User, error)
}
