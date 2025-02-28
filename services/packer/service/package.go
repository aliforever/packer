package service

import (
	"math"
	"sort"

	"packer/services/packer/models"
	"packer/services/packer/repository"
	"packer/services/packer/requests"
)

type Package struct {
	container repository.Container
}

func NewPackage(container repository.Container) *Package {
	return &Package{container: container}
}

func (p *Package) GetAll() ([]models.Package, error) {
	data, err := p.container.Packages().GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *Package) Add(req *requests.AddPackageRequest) (*models.Package, error) {
	pkg := &models.Package{
		ID:       req.ID,
		Quantity: req.Quantity,
	}

	if err := p.container.Packages().Upsert(pkg); err != nil {
		return nil, err
	}

	return pkg, nil
}

func (p *Package) RemoveByID(id string) error {
	return p.container.Packages().RemoveByID(id)
}

func (p *Package) CalculatePackages(quantity int) ([]*models.Package, error) {
	availablePackages, err := p.GetAll()
	if err != nil {
		return nil, err
	}

	return Distribute(availablePackages, quantity), nil
}

func Distribute(items []models.Package, quantity int) []*models.Package {
	itemsMap := map[string]*models.Package{}

	for _, item := range items {
		itemsMap[item.ID] = &item
	}

	minItemsOrdered := quantity

	sort.Slice(items, func(i, j int) bool {
		return items[i].Quantity < items[j].Quantity
	})

	minimumAmount := items[0].Quantity

	if minItemsOrdered%minimumAmount != 0 {
		minItemsOrdered = ((quantity / minimumAmount) + 1) * minimumAmount
	}

	dp := make([]float64, minItemsOrdered+1)

	for i := 1; i <= minItemsOrdered; i++ {
		dp[i] = math.Inf(1)
	}

	dp[0] = 0

	for i := 1; i <= minItemsOrdered; i++ {
		for _, packSize := range items {
			if i >= packSize.Quantity && dp[i-packSize.Quantity] != math.Inf(1) {
				dp[i] = math.Min(dp[i], dp[i-packSize.Quantity]+1)
			}
		}
	}

	result := make(map[string]int)

	remaining := minItemsOrdered

	for remaining > 0 {
		for _, packSize := range items {
			if remaining >= packSize.Quantity && dp[remaining] == dp[remaining-packSize.Quantity]+1 {
				result[packSize.ID] = result[packSize.ID] + 1
				remaining -= packSize.Quantity
				break
			}
		}
	}

	var finalResult []*models.Package

	for id, count := range result {
		for i := 0; i < count; i++ {
			finalResult = append(finalResult, itemsMap[id])
		}
	}

	return finalResult
}
