package middlewares

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/alireza-mf/go-ecommerce/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Input struct {
	Params any
	Body   any
	Query  any
}

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

func ValidateRequest[TInput any]() func(c *gin.Context) {
	return func(c *gin.Context) {
		// tinput := new(TInput)
		var tinput TInput
		input := (any)(tinput).(Input)
		params, _ := reflect.TypeOf(input).FieldByName("Params")
		body, _ := reflect.TypeOf(input).FieldByName("Body")
		query, _ := reflect.TypeOf(input).FieldByName("Query")

		// get the type of the struct
		//inputType := reflect.TypeOf(Input{})

		// >> params
		if params.Type.NumField() != 0 {
			for _, field := range reflect.VisibleFields(reflect.TypeOf(params)) {
				param := reflect.ValueOf(&params).Elem().FieldByName(field.Name)
				paramValue := c.Param(field.Tag.Get("json"))
				fmt.Println(param)
				fmt.Printf("%+v\n", field)
				fmt.Println(field.Name)
				fmt.Println(field.Tag.Get("json"), paramValue)

				// it can be either string or int, we should check if can be convertible to int
				if field.Type.Kind().String() == "string" {
					param.SetString(paramValue)
				} else {
					intParam, err := strconv.ParseInt(paramValue, 10, 64)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Param is not in number format."})
						return
					}
					param.SetInt(intParam)
				}
			}
		}

		fmt.Println(body.Type.NumField())
		// >> body
		if body.Type.NumField() != 0 {
			err := c.ShouldBindBodyWith(&input.Body, binding.JSON)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			fmt.Printf("%+v\n", reflect.ValueOf(&body))
		}

		// >> query
		if query.Type.NumField() != 0 {
			err := c.ShouldBindQuery(&query)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		// input = {{} {},  {}}
		fmt.Printf("%+v\n", input)
		c.Set("input", input)
		c.Next()
	}
}
