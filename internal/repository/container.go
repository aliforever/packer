package repository

type Container interface {
	Packages() Packages
}

type inMemoryContainer struct {
	packages Packages
}

func NewInMemoryContainer() Container {
	return inMemoryContainer{packages: NewInMemoryPackages()}
}

func (ic inMemoryContainer) Packages() Packages {
	return ic.packages
}
