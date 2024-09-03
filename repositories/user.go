package repositories

import (
	"context"
	"strings"

	"github.com/alireza-mf/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB         *mongo.Database
	Collection string
}

func ProvideUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{DB: db, Collection: "user"}
}

func (u *UserRepository) collection() *mongo.Collection {
	return u.DB.Collection(u.Collection)
}

// FindAll
func (u *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	cursor, err := u.collection().Find(context.TODO(), bson.D{})
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
func (u *UserRepository) FindByUserId(user_id string) (user *models.User, err error) {
	err = u.collection().FindOne(context.TODO(), bson.M{"user_id": user_id}).Decode(&user)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// FindByEmail
func (u *UserRepository) FindByEmail(email string) (user *models.User, err error) {
	err = u.collection().FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// Create
func (u *UserRepository) Create(userModel *models.User) (user *models.User, err error) {
	var id *mongo.InsertOneResult
	id, err = u.collection().InsertOne(context.TODO(), userModel)
	if err != nil {
		return nil, err
	}

	err = u.collection().FindOne(context.TODO(), bson.M{"_id": id.InsertedID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateByUserId
func (u *UserRepository) UpdateByUserId(user_id uint, updateKey string, updateValue interface{}) error {
	_, err := u.collection().UpdateOne(
		context.TODO(),
		bson.M{"user_id": user_id},
		bson.M{"$set": bson.M{updateKey: updateValue}},
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserId
func (u *UserRepository) DeleteByUserId(user_id uint) error {
	_, err := u.collection().DeleteOne(context.TODO(), bson.M{"user_id": user_id})

	if err != nil {
		return err
	}

	return nil
}
