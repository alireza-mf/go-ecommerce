package routers

import (
	"github.com/alireza-mf/go-ecommerce/controllers"
	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, controller controllers.UserController) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/user", middlewares.ValidateRequest[models.RegisterUserInput](), controller.RegisterUser)
			v1.POST("/user/login", middlewares.ValidateRequest[models.LoginUserInput](), controller.LoginUser)
			v1.GET("/user/:user_id", middlewares.Authorize(), middlewares.ValidateRequest[models.GetUserInput](), controller.GetUser)
		}
	}

	return r
}
