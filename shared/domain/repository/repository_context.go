package repository

type DB interface{}

// RepositoryContext 资源库上下文
type RepositoryContext interface {
	Instance() DB
	UnitOfWork
}
