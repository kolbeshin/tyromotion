package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"tyromotion/backend/internal/models"
	"tyromotion/backend/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func AddPatient(c *gin.Context, storage *postgres.Storage) {
	var patient models.Patient

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(rawBody, &patient)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	raw, err := storage.CreatePatient(patient)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, raw)
}
