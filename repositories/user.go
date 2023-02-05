package repositories

import (
	"context"
	"strings"

	"github.com/alireza-mf/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB *mongo.Database
	Collection string
}

func ProvideUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{DB: db, Collection: "user"}
}

// FindAll
func (u *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	cursor, err := u.DB.Collection(u.Collection).Find(context.TODO(), bson.D{})
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

// FindByUserId
func (u *UserRepository) FindByUserId(user_id uint) (user *models.User, err error) {
	err = u.DB.Collection(u.Collection).FindOne(context.TODO(), bson.M{"user_id": user_id}).Decode(&user)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// Create
func (u *UserRepository) Create(user *models.User) error {
	_, err := u.DB.Collection(u.Collection).InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}

// UpdateByUserId
func (u *UserRepository) UpdateByUserId(user_id uint, updateKey string, updateValue  interface{}) error {
	_, err := u.DB.Collection(u.Collection).UpdateOne(
		context.TODO(),
		bson.M{"user_id": user_id},
		bson.M{"$set": bson.M{updateKey: updateValue}},
	);

	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserId
func (u *UserRepository) DeleteByUserId(user_id uint) error {
	_, err := u.DB.Collection(u.Collection).DeleteOne(context.TODO(), bson.M{"user_id": user_id})
	
	if err != nil {
		return err
	}

	return nil
}
