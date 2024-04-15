package handlers

import (
	"net/http"
	"strconv"

	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func GetCompletedTreatmentsByPatient(c *gin.Context, storage *storage.Storage) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	raw, err := storage.GetCompletedTreatments(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, raw)
}
