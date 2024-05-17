package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"tyromotion/backend/internal/models"
	"tyromotion/backend/internal/storage/postgres"
)

func Register(c *gin.Context, storage *postgres.Storage) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:        data["name"],
		PhoneNumber: data["phone number"],
		Email:       data["email"],
		Password:    password,
	}

	raw, err := storage.RegisterUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, raw)
}
