package repository

// Repository 资源库基础接口
type Repository interface {
	Context() RepositoryContext
}
