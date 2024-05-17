package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

func ValidateToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("jwt")

	if err != nil {
		cookie = "Not Set"
		ctx.JSON(http.StatusUnauthorized, cookie)
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	ctx.Set("userID", claims.Issuer)
	ctx.Next()
}
