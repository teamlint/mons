package mysql

import (
	domain "github.com/teamlint/mons/shared/domain/repository"

	"github.com/jinzhu/gorm"
)

func DB(ctx domain.RepositoryContext) *gorm.DB {
	db, ok := ctx.Instance().(*gorm.DB)
	if !ok {
		panic("respository context insance convert to gorm.db failed")
	}
	return db
}
