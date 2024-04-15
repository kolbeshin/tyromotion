package srv

import (
	"net"

	"github.com/SenyashaGo/tyromotion/config"
	"github.com/SenyashaGo/tyromotion/srv/handlers"
	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func Run(storage *storage.Storage, config *config.Config) error {
	app := gin.Default()
	auth := app.Group("/auth")
	{
		auth.POST("/sign-in", func(c *gin.Context) {
			handlers.SignIn(c, storage)
		})
		auth.GET("/test", func(c *gin.Context) {
			handlers.Test(c, storage)
		})
		auth.POST("/test", func(c *gin.Context) {
			handlers.Test2(c, storage)
		})
	}

	patient := app.Group("/patients")
	{
		patient.GET("/all-patients", func(c *gin.Context) {
			handlers.GetAllPatients(c, storage)
		})
		patient.POST("/add-patient", func(c *gin.Context) {
			handlers.AddPatient(c, storage)
		})
		patient.GET("/completed-treatments/:id", func(c *gin.Context) {
			handlers.GetCompletedTreatmentsByPatient(c, storage)
		})
	}

	return app.Run(net.JoinHostPort(config.Host, config.Port))
}
