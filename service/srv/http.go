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
		auth.POST("/sign-in", func(c *gin.Context) {
			SignIn(c, storage)
		})
		auth.GET("/test", func(c *gin.Context) {
			Test(c, storage)
		})
		auth.POST("/test", func(c *gin.Context) {
			Test2(c, storage)
		})
	}

	return app.Run(net.JoinHostPort(config.Host, config.Port))
}
