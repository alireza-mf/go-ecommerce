package models

import "github.com/dgrijalva/jwt-go"

type ResponseSuccess struct {
	Data any `json:"data"`
}

type ResponseError struct {
	Error any `json:"error"`
}

type SortOrder int

const (
	Ascending  SortOrder = 1
	Descending SortOrder = -1
)

type JWTClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
