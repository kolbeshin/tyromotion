package handlers

import (
	"net/http"

	"github.com/SenyashaGo/tyromotion/storage"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context, storage *storage.Storage) {
	user, err := storage.GetDoctorByEmail("user2")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
