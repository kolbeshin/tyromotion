package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"tyromotion/backend/internal/models"
	"tyromotion/backend/internal/storage/postgres"
)

func Login(c *gin.Context, storage *postgres.Storage) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
	}

	user := models.User{
		Email:    data["email"],
		Password: []byte(data["password"]),
	}

	raw, err := storage.LoginUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)

	}

	if raw.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(raw.Password, []byte(data["password"])); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "wrong password"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(raw.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not login"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, raw)
}
