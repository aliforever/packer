package repository

import (
	"sync"

	"packer/internal/models"
)

type Packages interface {
	SeedDefault() error
	// GetAll returns all packages
	GetAll() ([]models.Package, error)
	// RemoveByID removes a package by ID
	RemoveByID(id string) error
	// Upsert upserts the model
	Upsert(packageModel *models.Package) error
}

// inMemoryPackages implements Packages interface
type inMemoryPackages struct {
	sync.Mutex

	packages []models.Package
}

// NewInMemoryPackages returns a new inMemoryPackages
func NewInMemoryPackages() Packages {
	return &inMemoryPackages{}
}

func (imp *inMemoryPackages) SeedDefault() error {
	defaultPackages := []models.Package{
		{
			ID:       "P_1",
			Quantity: 250,
		},
		{
			ID:       "P_2",
			Quantity: 500,
		},
		{
			ID:       "P_3",
			Quantity: 1000,
		},
		{
			ID:       "P_4",
			Quantity: 2000,
		},
		{
			ID:       "P_5",
			Quantity: 5000,
		},
	}

	availablePackages, err := imp.GetAll()
	if err != nil {
		return err
	}

	for _, p := range defaultPackages {
		found := false
		for _, ap := range availablePackages {
			if ap.Quantity == p.Quantity {
				found = true
				break
			}
		}

		if !found {
			err = imp.InsertPackage(&p)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetAll returns all packages
func (imp *inMemoryPackages) GetAll() ([]models.Package, error) {
	imp.Lock()
	defer imp.Unlock()

	return imp.packages, nil
}

// InsertPackage inserts a package
func (imp *inMemoryPackages) InsertPackage(packageModel *models.Package) error {
	imp.Lock()
	defer imp.Unlock()

	imp.packages = append(imp.packages, *packageModel)
	return nil
}

// RemoveByID removes a package by ID
func (imp *inMemoryPackages) RemoveByID(id string) error {
	imp.Lock()
	defer imp.Unlock()

	for i, p := range imp.packages {
		if p.ID == id {
			imp.packages = append(imp.packages[:i], imp.packages[i+1:]...)
			return nil
		}
	}

	return nil
}

// Upsert upserts the model
func (imp *inMemoryPackages) Upsert(pkg *models.Package) error {
	imp.Lock()
	defer imp.Unlock()

	for i, p := range imp.packages {
		if p.ID == pkg.ID {
			imp.packages[i] = *pkg
			return nil
		}
	}

	imp.packages = append(imp.packages, *pkg)

	return nil
}
