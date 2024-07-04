package app

import (
	"context"
	"tyromotion/backend/internal/api"
	"tyromotion/backend/internal/config"
	"tyromotion/backend/internal/storage/cache"
	"tyromotion/backend/internal/storage/postgres"
)

func Run(config *config.Config) error {
	newStorage, err := postgres.NewStorage(config.Dsn)
	if err != nil {
		return err
	}

	err = cache.NewRedisClient(config).Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	return api.Run(newStorage, config)
}
