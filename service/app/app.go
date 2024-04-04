package app

import (
	"github.com/SenyashaGo/tyromotion/config"
	"github.com/SenyashaGo/tyromotion/srv"
	"github.com/SenyashaGo/tyromotion/storage"
)

func Run(config *config.Config) error {
	newStorage, err := storage.NewStorage(config.Dsn)
	if err != nil {
		return err
	}

	return srv.Run(newStorage, config)
}
