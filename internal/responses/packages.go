package responses

import (
	"packer/internal/models"
)

type Packages struct {
	Packages []models.Package `json:"packages,omitempty"`
}

type CalculatedPackages struct {
	Packages []*models.Package `json:"packages,omitempty"`
}
