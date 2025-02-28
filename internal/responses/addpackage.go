package responses

import (
	"packer/internal/models"
)

type AddPackage struct {
	Package *models.Package `json:"package,omitempty"`
}
