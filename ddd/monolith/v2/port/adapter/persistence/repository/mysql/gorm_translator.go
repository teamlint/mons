package mysql

import (
	"github.com/jinzhu/gorm"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

func DB(ctx shared.RepositoryContext) *gorm.DB {
	db, ok := ctx.Instance().(*gorm.DB)
	if !ok {
		panic("respository context insance convert to gorm.db failed")
	}
	return db
}
