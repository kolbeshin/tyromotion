package app

import (
	"tyromotion/backend/internal/api"
	"tyromotion/backend/internal/config"
	"tyromotion/backend/internal/storage/postgres"
)

func Run(config *config.Config) error {
	newStorage, err := postgres.NewStorage(config.Dsn)
	if err != nil {
		return err
	}

	return api.Run(newStorage, config)
}
