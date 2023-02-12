package db

import (
	"context"

	"github.com/alireza-mf/go-ecommerce/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Initialize MongoDB connection.
func InitMongoDB() (*mongo.Database, error) {

	// Make MongoDB Connection
	clientOption := options.Client().ApplyURI(config.GetConfig().MongoURI)
	db, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}

	// Check Connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return db.Database(config.GetConfig().MongoDatabaseName), nil
}
