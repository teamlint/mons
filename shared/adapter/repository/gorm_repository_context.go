package repository

import (
	"log"

	domain "github.com/teamlint/mons/shared/domain/repository"

	"github.com/jinzhu/gorm"
)

// GormRepositoryContext GORM 资源库上下文实现
type GormRepositoryContext struct {
	db *gorm.DB
}

func NewGormRepositoryContext(db *gorm.DB) domain.RepositoryContext {
	return &GormRepositoryContext{db: db}
}

func (c *GormRepositoryContext) Begin() (domain.RepositoryContext, error) {
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

func (c *GormRepositoryContext) Instance() domain.DB {
	return c.db
}
