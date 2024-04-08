package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SenyashaGo/tyromotion/models"
	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func AddPatient(c *gin.Context, storage *storage.Storage) {
	var patient models.Patient

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(rawBody, &patient)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	raw, err := storage.CreatePatient(patient)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, raw)
}
