package routers

import (
	"github.com/alireza-mf/go-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, controller controllers.UserController) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/user/:user_id", controller.GetUser)
		}
	}

	return r
}
