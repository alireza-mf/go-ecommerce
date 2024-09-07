package controllers

import (
	"net/http"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/services"
	"github.com/alireza-mf/go-ecommerce/util"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.ProductService
}

func ProvideProductController(u services.ProductService) ProductController {
	return ProductController{ProductService: u}
}

func (u *ProductController) CreateProduct(c *gin.Context) {
	input := c.MustGet("input").(models.CreateProductInput)

	response, err := u.ProductService.CreateProduct(&input.Body)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError)
		return
	}

	util.ResponseSuccess(c, &response)
}

func (u *ProductController) GetProduct(c *gin.Context) {
	input := c.MustGet("input").(models.GetProductInput)

	product, err := u.ProductService.FindById(input.Params.ProductId)

	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError)
		return
	}
	if product == nil {
		util.ResponseError(c, http.StatusNotFound)
		return
	}

	util.ResponseSuccess(c, &product)
}

func (u *ProductController) GetProducts(c *gin.Context) {
	// input := c.MustGet("input").(models.GetProductsInput)

	products, err := u.ProductService.FindAll()

	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError)
		return
	}

	util.ResponseSuccess(c, products)
}
