package models

import "time"

// User represents User
type User struct {
	UserId          string    `bson:"user_id" json:"user_id"`
	Email           string    `bson:"email" json:"email"`
	Password        string    `bson:"password" json:"-"`
	DeliveryAddress *string   `bson:"delivery_address" json:"delivery_address"`
	CreatedAt       time.Time `bson:"created_at" json:"created_at"`
}

type UserInputForm struct {
	Email           string  `json:"email" binding:"required,email"`
	Password        string  `json:"password" binding:"required,min=6"`
	DeliveryAddress *string `json:"delivery_address"`
}

type UserLoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
