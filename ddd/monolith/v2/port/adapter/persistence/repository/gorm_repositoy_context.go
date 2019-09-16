package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

// GormRepositoryContext GORM 资源库上下文实现
type GormRepositoryContext struct {
	db *gorm.DB
}

func NewGormRepositoryContext(db *gorm.DB) shared.RepositoryContext {
	return &GormRepositoryContext{db: db}
}

func (c *GormRepositoryContext) Begin() (shared.RepositoryContext, error) {
	log.Println("[GormRepositoryContext] begin")
	tx := c.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &GormRepositoryContext{db: tx}, nil
}
func (c *GormRepositoryContext) Commit() error {
	log.Println("[GormRepositoryContext] commit")
	return c.db.Commit().Error
}

func (c *GormRepositoryContext) Rollback() error {
	log.Println("[GormRepositoryContext] rollback")
	return c.db.Rollback().Error
}

func (c *GormRepositoryContext) Instance() shared.DB {
	return c.db
}
