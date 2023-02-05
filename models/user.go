package models

// User represents User
type User struct {
	UserId          uint32 `bson:"user_id" json:"user_id"`
	Email           string `bson:"email" json:"email"`
	Password        string `bson:"password" json:"password"`
	DeliveryAddress string `bson:"delivery_address" json:"delivery_address"`
}
