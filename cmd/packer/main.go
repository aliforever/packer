package main

import (
	"log/slog"
	"net/http"
	"os"

	"packer/internal/config"
	"packer/internal/controller"
	"packer/internal/repository"
	"packer/internal/routes"
	"packer/internal/service"
)

func main() {
	cfg, err := config.ParseFromEnv()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(cfg.LogLevel),
	}))

	inMemoryDbContainer := repository.NewInMemoryContainer()

	if cfg.SeedDefault {
		err = inMemoryDbContainer.Packages().SeedDefault()
		if err != nil {
			logger.Error("couldn't seed default packages")
			return
		}
	}

	svc := service.NewPackage(inMemoryDbContainer)

	ctr := controller.NewPackage(svc, logger)

	mux := http.NewServeMux()

	routes.Init(mux, ctr)

	logger.Info("starting server", slog.String("address", cfg.HttpAddress))

	err = http.ListenAndServe(cfg.HttpAddress, mux)
	if err != nil {
		logger.Error("error starting server", slog.Any("error", err))
	}
}
