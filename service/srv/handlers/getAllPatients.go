package handlers

import (
	"net/http"

	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context, storage *storage.Storage) {
	patients, err := storage.GetAllPatientsFromTable()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, patients)
}
