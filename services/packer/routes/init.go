package routes

import (
	"net/http"

	"packer/services/packer/controller"
)

func Init(mux *http.ServeMux, packageController *controller.Package) {
	mux.HandleFunc("/packages", packageController.GetAll)
	mux.HandleFunc("/packages/add", packageController.Add)
	mux.HandleFunc("/packages/remove", packageController.RemoveByID)
	mux.HandleFunc("/packages/order", packageController.CalculatePackages)

	mux.Handle("/", http.FileServer(http.Dir("public")))
}
