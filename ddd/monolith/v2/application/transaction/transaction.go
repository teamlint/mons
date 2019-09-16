package transaction

import (
	"fmt"

	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

// Scope 开启事务处理范围
func Scope(ctx shared.RepositoryContext, fn func(shared.RepositoryContext) error) (err error) {
	// db.Tx
	tx, err := ctx.Begin()
	if err != nil {
		return
	}
	// unit of work
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
			// panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(tx)
	return err
}
