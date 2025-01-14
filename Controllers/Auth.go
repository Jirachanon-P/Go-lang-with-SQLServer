package Controllers

import (
	"first-api/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserLogin struct {
	Username string `json:"username" binding:"require"`
	Password string `json:"password" binding:"require"`
}

// Login godoc
// @summary Login
// @description Login
// @tags Authentication
// @id Login
// @Param LoginForm body UserLogin true "Username & Password"
// @accept json
// @produce json
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	passowrd := user.Password
	err := Models.Login(&user, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Models.Response{
			Code:   404,
			Status: "Not Found",
			Data:   err.Error(),
		})
		return
	}

	if passowrd != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Models.Response{
			Code:   401,
			Status: "unauthorized",
			Data:   "wrong passoword!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(5 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte("Token_Secret_Key"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	c.JSON(http.StatusOK, Models.Response{
		Code:   200,
		Status: "Login Success!",
		Data:   ss,
	})
}

func GetUserFromToken(c *gin.Context) (id int) {
	claims, exists := c.Get("claims")
	if !exists {
		return 0
	}
	if jwtClaims, ok := claims.(jwt.MapClaims); ok {
		userID := jwtClaims["id"]

		return int(userID.(float64))
	} else {
		return -1
	}
}
