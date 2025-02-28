package responses

import "packer/services/packer/models"

type AddPackage struct {
	Package *models.Package `json:"package,omitempty"`
}
