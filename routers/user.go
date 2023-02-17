package routers

import (
	"github.com/alireza-mf/go-ecommerce/controllers"
	"github.com/alireza-mf/go-ecommerce/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, controller controllers.UserController) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/user", controller.RegisterUser)
			v1.POST("/user/login", controller.LoginUser)
			v1.Use(middlewares.Authorize()).GET("/user/:user_id", controller.GetUser)
		}
	}

	return r
}
