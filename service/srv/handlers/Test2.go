package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SenyashaGo/tyromotion/models"
	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func Test2(c *gin.Context, storage *storage.Storage) {
	var doctor models.Doctor

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(rawBody, &doctor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	raw, err := storage.CreateDoctor(doctor)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, raw)
}
