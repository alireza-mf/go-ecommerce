package repositories

import (
	"context"
	"log"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Entity string

const (
	Product               Entity = "product"
	ProductAttribute      Entity = "product_attribute"
	ProductAttributeValue Entity = "product_attribute_value"
)

type ProductRepository struct {
	DB *mongo.Database
}

func ProvideProductRepository(db *mongo.Database) ProductRepository {
	return ProductRepository{DB: db}
}

func (r *ProductRepository) collection(name Entity) *mongo.Collection {
	return r.DB.Collection(string(name))
}

// FindAll
func (r *ProductRepository) FindAll() ([]models.Product, error) {
	// Aggregation pipeline
	pipeline := mongo.Pipeline{
		{
			{"$lookup", bson.D{
				{"from", ProductAttributeValue},
				{"localField", "product_id"},
				{"foreignField", "product_id"},
				{"as", "attribute_values"},
			}},
		},
		{
			{"$lookup", bson.D{
				{"from", ProductAttribute},
				{"localField", "attribute_values.product_attribute_id"},
				{"foreignField", "product_attribute_id"},
				{"as", "attributes"},
			}},
		},
	}

	cursor, err := r.collection(Product).Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Panic(err)
		return nil, err
	}

	return products, nil
}

// FindByProductId
func (r *ProductRepository) FindByProductId(productId string) (product *models.Product, err error) {
	// Aggregation pipeline
	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{{"product_id", productId}}},
		},
		{
			{"$lookup", bson.D{
				{"from", ProductAttributeValue},
				{"localField", "product_id"},
				{"foreignField", "product_id"},
				{"as", "attribute_values"},
			}},
		},
		{
			{"$lookup", bson.D{
				{"from", ProductAttribute},
				{"localField", "attribute_values.product_attribute_id"},
				{"foreignField", "product_attribute_id"},
				{"as", "attributes"},
			}},
		},
	}

	cursor, err := r.collection(Product).Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Panic(err)
		return nil, err
	}

	if len(products) == 0 {
		return nil, nil
	}

	log.Print(&products[0])
	return &products[0], nil
}

// Create
func (r *ProductRepository) Create(productModel models.Product) (product *models.Product, err error) {
	productsCollection := r.collection(Product)

	var id *mongo.InsertOneResult

	// Insert the product
	id, err = productsCollection.InsertOne(context.TODO(), productModel)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	// Insert attribute values
	if len(productModel.AttributeValues) != 0 {
		attributeValuesCollection := r.collection(ProductAttributeValue)
		var attributeValues []interface{}

		for _, value := range productModel.AttributeValues {
			attributeValues = append(attributeValues, value)
		}

		_, err = attributeValuesCollection.InsertMany(context.TODO(), attributeValues)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
	}

	err = r.collection(Product).FindOne(context.TODO(), bson.M{"_id": id.InsertedID}).Decode(&product)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return product, nil
}

// UpdateByProductId
func (r *ProductRepository) UpdateByProductId(product_id uint, productModel *models.Product) error {
	productBson, _ := util.ToBsonM(productModel)
	_, err := r.collection(Product).UpdateOne(
		context.TODO(),
		bson.M{"product_id": product_id},
		bson.M{"$set": productBson},
	)

	if err != nil {
		return err
	}

	for _, value := range productModel.AttributeValues {
		valueBson, _ := util.ToBsonM(value)
		_, err = r.collection(ProductAttributeValue).UpdateOne(
			context.TODO(),
			bson.M{"product_attribute_value_id": value.ProductAttributeValueId, "product_id": product_id},
			bson.M{"$set": valueBson},
		)
		if err != nil {
			return err
		}

	}

	return nil
}

// DeleteByProductId
func (r *ProductRepository) DeleteByProductId(product_id uint) error {
	_, err := r.collection(Product).DeleteOne(context.TODO(), bson.M{"product_id": product_id})
	if err != nil {
		return err
	}

	_, err = r.collection(ProductAttributeValue).DeleteMany(context.TODO(), bson.M{"product_id": product_id})
	if err != nil {
		return err
	}

	return nil
}
