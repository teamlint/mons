package repository

// UnitOfWork 工作单元处理接口
type UnitOfWork interface {
	Begin() (RepositoryContext, error)
	Commit() error
	Rollback() error
}
