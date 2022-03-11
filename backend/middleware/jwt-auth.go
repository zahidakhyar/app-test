package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zahidakhyar/app-test/backend/helper"
	auth_service "github.com/zahidakhyar/app-test/backend/src/auth/service"
)

func AuthorizeJwt(jwtService auth_service.JwtServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "Authorization header is missing", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if !token.Valid {
			log.Println("Token is invalid: ", err)
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
