package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	LogLevel int

	HttpAddress string

	SeedDefault bool
}

func ParseFromEnv() (*Config, error) {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		return nil, errors.New("LOG_LEVEL is required")
	}

	logLevelInt, err := strconv.Atoi(logLevel)
	if err != nil {
		return nil, errors.New("invalid log level")
	}

	httpAddress := os.Getenv("HTTP_ADDRESS")
	if httpAddress == "" {
		return nil, errors.New("HTTP_ADDRESS is required")
	}

	seedDefault := os.Getenv("SEED_DEFAULT")
	if seedDefault == "" {
		return nil, errors.New("SEED_DEFAULT is required")
	}

	seedDefaultBool, err := strconv.ParseBool(seedDefault)
	if err != nil {
		return nil, errors.New("invalid seed default value")
	}

	return &Config{
		LogLevel:    logLevelInt,
		HttpAddress: httpAddress,
		SeedDefault: seedDefaultBool,
	}, nil
}
