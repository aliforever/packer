package models

type Package struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
	NumPacks int    `json:"numPacks,omitempty"`
}
