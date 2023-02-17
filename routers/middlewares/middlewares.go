package middlewares

import (
	"net/http"
	"strings"

	"github.com/alireza-mf/go-ecommerce/util/jwt"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token."})
			context.Abort()
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if tokenParts[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token format is wrong. 'Bearer' is not contained."})
			context.Abort()
			return
		}

		err := jwt.ValidateToken(tokenParts[1])
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
