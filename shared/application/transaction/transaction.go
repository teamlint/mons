package transaction

import (
	"fmt"

	domain "github.com/teamlint/mons/shared/domain/repository"
)

// Scope begin transaction scope
func Scope(ctx domain.RepositoryContext, fn func(domain.RepositoryContext) error) (err error) {
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
