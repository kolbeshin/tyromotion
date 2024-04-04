package srv

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SenyashaGo/tyromotion/models"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var doctor models.Doctor

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(rawBody, &doctor)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Status(http.StatusNoContent)
}
