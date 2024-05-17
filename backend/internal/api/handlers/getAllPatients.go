package handlers

import (
	"net/http"
	"tyromotion/backend/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context, storage *postgres.Storage) {
	patients, err := storage.GetAllPatientsFromTable()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, patients)
}
