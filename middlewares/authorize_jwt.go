package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/Rickykn/buddyku-app.git/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("SECRETKEY")), nil
	})

}

func AuthorizeJWTAdmin(c *gin.Context) {

	var authHeader = ""

	authHeader = c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unathorize",
		})
	}

	tokenString := authHeader[7:]

	token, err := validateToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unathorize",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unathorize",
			},
		)
	}

	userJson, _ := json.Marshal(claims["admin"])
	admin := models.Admin{}
	err = json.Unmarshal(userJson, &admin)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unathorize",
			},
		)
	}

	c.Set("admin", admin)

}

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unathorize",
		})
	}

	tokenString := authHeader[7:]

	token, err := validateToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unathorize",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unathorize",
			},
		)
	}

	userJson, _ := json.Marshal(claims["user"])
	user := models.User{}
	err = json.Unmarshal(userJson, &user)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unathorize",
			},
		)
	}

	c.Set("user", user)

}
