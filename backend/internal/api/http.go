package api

import (
	"github.com/gin-gonic/gin"
	"net"
	"tyromotion/backend/internal/api/handlers"
	"tyromotion/backend/internal/config"
	"tyromotion/backend/internal/middlewares"
	"tyromotion/backend/internal/storage/postgres"
)

func Run(storage *postgres.Storage, config *config.Config) error {
	app := gin.Default()
	auth := app.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			handlers.Login(c, storage)
		})
		auth.GET("/logout", func(c *gin.Context) {
			handlers.Logout(c)
		})
		auth.POST("/register", func(c *gin.Context) {
			handlers.Register(c, storage)
		})
	}

	patient := auth.Group("/patients")
	patient.Use(middlewares.ValidateToken)
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
