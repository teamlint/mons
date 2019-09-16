package repository

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
)

type GormTransaction struct {
	db *gorm.DB
}

func NewGormTransaction(db *gorm.DB) shared.TransactionScoper {
	return &GormTransaction{db: db}
}
func (t *GormTransaction) Scope(fn func(shared.RepositoryContext) error) error {
	return transactionScope(t.db, fn)
}

// func transactionScope(db *gorm.DB, fn func(*sql.Tx) error) (err error) {
func transactionScope(db *gorm.DB, fn func(shared.RepositoryContext) error) (err error) {
	// gorm.Tx
	gormTx := db.Begin()
	if gormTx.Error != nil {
		err = gormTx.Error
		return
	}
	// sql.Tx
	tx, ok := gormTx.CommonDB().(*sql.Tx)
	if !ok {
		return fmt.Errorf("gorm db convert to sql.Tx error")
	}
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
	rc := NewGormRepositoryContext(gormTx)
	err = fn(rc)
	return err
}
