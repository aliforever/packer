package requests

type AddPackageRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}
