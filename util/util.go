package util

import (
	"net/http"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// If is a simple one-line if statement
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func ToBsonM(data interface{}) (bson.M, error) {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	var bsonMap bson.M
	err = bson.Unmarshal(bsonBytes, &bsonMap)
	if err != nil {
		return nil, err
	}
	return bsonMap, nil
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

func GetUserClaims(c *gin.Context) (user *models.JWTClaim, exists bool) {
	value, ok := c.Get("user")
	if !exists {
		return nil, false
	}

	claims, ok := value.(*models.JWTClaim)
	if !ok || claims == nil {
		return nil, false
	}

	return claims, true
}
