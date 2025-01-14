package models

import (
	"time"
)

// User represents Product
type Product struct {
	ProductId       string                  `bson:"product_id" json:"product_id" uri:"product_id"`
	Name            string                  `bson:"name" json:"name"`
	Price           uint                    `bson:"price" json:"price"`
	IsActive        bool                    `bson:"is_active" json:"is_active"`
	AttributeValues []ProductAttributeValue `bson:"-" json:"attribute_value"`
	CreatedAt       time.Time               `bson:"created_at" json:"created_at"`
}

// User represents ProductAttribute
type ProductAttribute struct {
	ProductAttributeId string                  `bson:"product_attribute_id" json:"product_attribute_id"`
	Name               string                  `bson:"name" json:"name"`
	Values             []ProductAttributeValue `bson:"-" json:"values"`
	CreatedAt          time.Time               `bson:"created_at" json:"created_at"`
}

// User represents ProductAttributeValue
type ProductAttributeValue struct {
	ProductAttributeValueId string    `bson:"product_attribute_value_id" json:"product_attribute_value_id"`
	ProductId               string    `bson:"product_id" json:"product_id"`
	ProductAttributeId      string    `bson:"product_attribute_id" json:"product_attribute_id"`
	Value                   string    `bson:"value" json:"value"`
	CreatedAt               time.Time `bson:"created_at" json:"created_at"`
}

type CreateProductModel struct {
	Name            string `json:"name" binding:"required,min=2"`
	Price           uint   `json:"price" binding:"required"`
	AttributeValues []struct {
		ProductAttributeId string `json:"product_attribute_id"`
		Value              string `json:"value" binding:"required"`
	} `json:"attribute_values"`
	IsActive *bool `json:"is_active"`
}

// ### Request Inputs ###

type CreateProductInput struct {
	Params struct{}
	Body   CreateProductModel
	Query  struct{}
}

type GetProductInput struct {
	Params struct {
		ProductId string `uri:"product_id"`
	}
	Body  struct{}
	Query struct{}
}

type GetProductsInput struct {
	Params struct{}
	Body   struct{}
	Query  struct {
		IsActive  *bool             `form:"is_active"`
		PriceFrom *uint             `form:"price_from"`
		PriceTo   *uint             `form:"price_to"`
		SortField *ProductSortField `form:"sort_field" binding:"omitempty,oneof=name price created_at"`
		SortOrder *SortOrder        `form:"sort_order" binding:"omitempty,oneof=1 -1"`
	}
}

// ### Others ###

type ProductSortField string

const (
	SortProductByName      ProductSortField = "name"
	SortProductByPrice     ProductSortField = "price"
	SortProductByCreatedAt ProductSortField = "created_at"
)

type ProductFilterOptions struct {
	IsActive  *bool
	PriceFrom *uint
	PriceTo   *uint
	SortField *ProductSortField
	SortOrder *SortOrder
}
