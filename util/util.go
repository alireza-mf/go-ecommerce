package util

import (
	"net/http"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/gin-gonic/gin"
)

// If is a simple one-line if statement
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func ResponseSuccess[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, models.ResponseSuccess{Data: data})
}

func ResponseError(c *gin.Context, code int, message ...string) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	c.JSON(code, models.ResponseError{Error: message})
}
