package http

import (
	"fmt"
	"net/http"

	"github.com/MuhammadSuryono1997/framework-okta/base/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	fmt.Println("Inside auth jwt middleware")

	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		fmt.Println("Auth Header = ", authHeader)

		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NOT_FOUND.AsInvalidResponse())
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			// c.JSON(http.StatusOK, )
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
