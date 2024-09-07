package services

import (
	"time"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/repositories"
	"github.com/google/uuid"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func ProvideProductService(t repositories.ProductRepository) ProductService {
	return ProductService{ProductRepository: t}
}

// FindById
func (u *ProductService) FindById(ProductId string) (*models.Product, error) {
	return u.ProductRepository.FindByProductId(ProductId)
}

// FindById
func (u *ProductService) FindAll() ([]models.Product, error) {
	return u.ProductRepository.FindAll()
}

// CreateProduct
func (u *ProductService) CreateProduct(model *models.CreateProductModel) (*models.Product, error) {

	productModel := models.Product{
		ProductId:       uuid.Must(uuid.NewRandom()).String(),
		Name:            model.Name,
		Price:           model.Price,
		CreatedAt:       time.Now(),
		AttributeValues: []models.ProductAttributeValue{},
	}

	if model.IsActive == nil {
		productModel.IsActive = true
	} else {
		productModel.IsActive = *model.IsActive
	}

	for _, value := range model.AttributeValues {
		// check if attribute exists
		productModel.AttributeValues = append(productModel.AttributeValues, models.ProductAttributeValue{
			ProductAttributeValueId: uuid.Must(uuid.NewRandom()).String(),
			ProductId:               productModel.ProductId,
			ProductAttributeId:      value.ProductAttributeId,
			Value:                   value.Value,
			CreatedAt:               time.Now(),
		})
	}

	return u.ProductRepository.Create(productModel)
}
