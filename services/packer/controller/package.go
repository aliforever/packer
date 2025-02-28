package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"packer/services/packer/requests"
	"packer/services/packer/responses"
	"packer/services/packer/service"
)

type Package struct {
	service *service.Package

	logger *slog.Logger
}

func NewPackage(service *service.Package, logger *slog.Logger) *Package {
	return &Package{service: service, logger: logger}
}

func (p *Package) GetAll(wr http.ResponseWriter, _ *http.Request) {
	data, err := p.service.GetAll()
	if err != nil {
		p.logger.Error("error getting packages", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(responses.Ok(responses.Packages{Packages: data}))
	if err != nil {
		p.logger.Error("error marshalling response", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	wr.Header().Set("Content-Type", "application/json")
	_, _ = wr.Write(j)

	return
}

func (p *Package) Add(wr http.ResponseWriter, req *http.Request) {
	var reqBody requests.AddPackageRequest

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		p.logger.Error("error decoding request body", slog.Any("error", err))
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	pkg, err := p.service.Add(&reqBody)
	if err != nil {
		p.logger.Error("error adding package", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(responses.Ok(responses.AddPackage{Package: pkg}))
	if err != nil {
		p.logger.Error("error marshalling response", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	wr.Header().Set("Content-Type", "application/json")
	_, _ = wr.Write(j)
}

func (p *Package) RemoveByID(wr http.ResponseWriter, req *http.Request) {
	var reqBody requests.RemovePackageRequest

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		p.logger.Error("error decoding request body", slog.Any("error", err))
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.service.RemoveByID(reqBody.ID)
	if err != nil {
		p.logger.Error("error removing package", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(responses.Ok(nil))
	if err != nil {
		p.logger.Error("error marshalling response", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	wr.Header().Set("Content-Type", "application/json")
	_, _ = wr.Write(j)
}

func (p *Package) CalculatePackages(wr http.ResponseWriter, req *http.Request) {
	var reqBody requests.SubmitOrderRequest

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		p.logger.Error("error decoding request body", slog.Any("error", err))
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	packages, err := p.service.CalculatePackages(reqBody.Quantity)
	if err != nil {
		p.logger.Error("error calculating packages", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(responses.Ok(responses.CalculatedPackages{Packages: packages}))
	if err != nil {
		p.logger.Error("error marshalling response", slog.Any("error", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	wr.Header().Set("Content-Type", "application/json")
	_, _ = wr.Write(j)
}
