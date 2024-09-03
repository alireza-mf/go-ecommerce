package middlewares

import (
	"net/http"
	"reflect"
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

// ValidateRequest is a middleware for validating input based on the given generic type.
func ValidateRequest[TInput any]() func(c *gin.Context) {
	return func(c *gin.Context) {
		var input TInput

		val := reflect.ValueOf(&input).Elem()

		// params
		paramsField := val.FieldByName("Params")
		if paramsField.IsValid() && paramsField.CanSet() {
			if err := c.ShouldBindUri(paramsField.Addr().Interface()); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		// body
		bodyField := val.FieldByName("Body")
		if bodyField.IsValid() && bodyField.CanSet() {
			if err := c.ShouldBindJSON(bodyField.Addr().Interface()); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		// query
		queryField := val.FieldByName("Query")
		if queryField.IsValid() && queryField.CanSet() {
			if err := c.ShouldBindQuery(queryField.Addr().Interface()); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		// fmt.Printf("%+v\n", input)
		c.Set("input", input)
		c.Next()
	}
}
