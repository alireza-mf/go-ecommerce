package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Initialize MongoDB connection.
func InitMongoDB() (*mongo.Client, error) {

	// TODO: get URI form env
	// Make MongoDB Connection
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}

	// Check Connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}