package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v11"
)

func New() (*Service, error) {
	var envValues envValues

	err := env.Parse(&envValues)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	var storageFolder string

	if envValues.StorageFolder == "" {
		var err error

		storageFolder, err = userAppStorage()
		if err != nil {
			return nil, err
		}
	} else {
		storageFolder = envValues.StorageFolder
	}

	return &Service{
		envValues:     envValues,
		storageFolder: storageFolder,
	}, nil
}

type Service struct {
	envValues     envValues
	storageFolder string
}

type envValues struct {
	StorageFolder string `env:"STORAGE_FOLDER"`
}

func (s *Service) StorageFolder() string {
	return s.storageFolder
}

func userAppStorage() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(homeDir, "AppData", "pover"), nil
	case "linux":
		return filepath.Join(homeDir, ".local", "share", "pover"), nil
	case "darwin":
		return filepath.Join(homeDir, "Library", "Application Support", "pover"), nil
	default:
		return filepath.Join(homeDir, ".pover"), nil
	}
}
