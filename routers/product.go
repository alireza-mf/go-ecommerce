package routers

import (
	"github.com/alireza-mf/go-ecommerce/controllers"
	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.Engine, controller controllers.ProductController) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/product", middlewares.Authorize(), middlewares.UserRole("admin"), middlewares.ValidateRequest[models.CreateProductInput](), controller.CreateProduct)
			v1.GET("/product/:product_id", middlewares.ValidateRequest[models.GetProductInput](), controller.GetProduct)
			v1.GET("/product", middlewares.ValidateRequest[models.GetProductsInput](), controller.GetProducts)
		}
	}

	return r
}
