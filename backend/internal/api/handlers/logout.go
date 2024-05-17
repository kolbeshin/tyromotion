package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "success")
}
