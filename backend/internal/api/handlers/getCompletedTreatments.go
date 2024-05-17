package handlers

import (
	"net/http"
	"strconv"
	"tyromotion/backend/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func GetCompletedTreatmentsByPatient(c *gin.Context, storage *postgres.Storage) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	raw, err := storage.GetCompletedTreatments(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, raw)
}
