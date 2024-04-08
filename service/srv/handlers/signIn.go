package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SenyashaGo/tyromotion/models"
	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context, storage *storage.Storage) {
	var doctor models.Doctor

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(rawBody, &doctor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	user, err := storage.GetDoctorByEmail(doctor.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if user.Password != doctor.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
