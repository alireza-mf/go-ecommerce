package jwt

import (
	"errors"
	"time"

	"github.com/alireza-mf/go-ecommerce/config"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.GetConfig().JWTSecret)

type JWTClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, userId string) (tokenString string, err error) {
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &JWTClaim{
		Email:  email,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("Couldn't parse claims.")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired.")
		return
	}
	return
}
