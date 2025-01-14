package Routes

import (
	"first-api/Models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func authorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if token1, err := validateToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Models.Response{
			Code:   401,
			Status: "Unauthorized",
			Data:   nil,
		})
		return
	} else {
		if claims, ok := token1.Claims.(jwt.MapClaims); ok && token1.Valid {
			c.Set("claims", claims)
		}
	}
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("Token_Secret_Key"), nil
	})
}
