package srv

import (
	"net"

	"github.com/SenyashaGo/tyromotion/config"
	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func Run(storage *storage.Storage, config *config.Config) error {
	app := gin.Default()

	auth := app.Group("/auth")
	{
		auth.POST("/sign-in", SignIn)
	}

	return app.Run(net.JoinHostPort(config.Host, config.Port))
}
