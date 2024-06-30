package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

func New() (*Service, error) {
	var envValues envValues

	err := env.Parse(&envValues)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	return &Service{
		envValues: envValues,
	}, nil
}

type Service struct {
	envValues envValues
}

type envValues struct {
	StorageFolder string `env:"STORAGE_FOLDER,required"`
}

func (s *Service) StorageFolder() string {
	return s.envValues.StorageFolder
}
