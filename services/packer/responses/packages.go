package responses

import "packer/services/packer/models"

type Packages struct {
	Packages []models.Package `json:"packages,omitempty"`
}

type CalculatedPackages struct {
	Packages []*models.Package `json:"packages,omitempty"`
}
