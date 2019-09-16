package container

import "go.uber.org/dig"

type Container struct {
	container *dig.Container
}

// ProviderOption a provideOption modifies the default behavior of Provide.
type ProvideOption interface {
	// applyProvideOption(*provideOptions)
	// applyProvideOption()
}
type InvokeOption interface {
	// unimplemented()
}

func New() *Container {
	return &Container{container: dig.New()}
}
func (c *Container) Provide(ctor interface{}, opts ...ProvideOption) error {
	return c.container.Provide(ctor)
}

func (c *Container) Invoke(fn interface{}, opts ...InvokeOption) error {
	return c.container.Invoke(fn)
}
