package services

import (
	"errors"
	"time"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/repositories"
	"github.com/alireza-mf/go-ecommerce/util"
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

// FindAll
func (u *ProductService) FindAll(sortOptions *models.ProductFilterOptions) ([]models.Product, error) {
	defaultSortField := models.SortProductByCreatedAt
	defaultSortOrder := models.Descending

	sortOptions = &models.ProductFilterOptions{
		IsActive:  sortOptions.IsActive,
		PriceFrom: sortOptions.PriceFrom,
		PriceTo:   sortOptions.PriceTo,
		SortField: util.If(sortOptions.SortField != nil, sortOptions.SortField, &defaultSortField),
		SortOrder: util.If(sortOptions.SortOrder != nil, sortOptions.SortOrder, &defaultSortOrder),
	}

	if sortOptions.PriceFrom != nil && sortOptions.PriceTo != nil && *sortOptions.PriceFrom > *sortOptions.PriceTo {
		return nil, errors.New("FindAll::price_from_price_to")
	}

	return u.ProductRepository.FindAll(sortOptions)
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
