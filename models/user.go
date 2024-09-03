package models

import (
	"time"
)

// User represents User
type User struct {
	UserId          string    `bson:"user_id" json:"user_id" uri:"user_id"`
	Email           string    `bson:"email" json:"email"`
	Password        string    `bson:"password" json:"-"`
	DeliveryAddress *string   `bson:"delivery_address" json:"delivery_address"`
	CreatedAt       time.Time `bson:"created_at" json:"created_at"`
}

type RegisterUser struct {
	Email           string  `json:"email" binding:"required,email"`
	Password        string  `json:"password" binding:"required,min=6"`
	DeliveryAddress *string `json:"delivery_address"`
}

// ### Request Inputs ###

type RegisterUserInput struct {
	Params struct{}
	Body   RegisterUser
	Query  struct{}
}

type LoginUserInput struct {
	Params struct{}
	Body   struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	Query struct{}
}

type GetUserInput struct {
	Params struct {
		UserId string `bson:"user_id" json:"user_id" uri:"user_id"`
	}
	Body  struct{}
	Query struct{}
}
