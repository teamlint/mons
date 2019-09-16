package repository

// TransactionScoper 事务范围接口
type TransactionScoper interface {
	Scope(fn func(RepositoryContext) error) error
}
