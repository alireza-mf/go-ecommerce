package jwt

import (
	"errors"
	"time"

	"github.com/alireza-mf/go-ecommerce/config"
	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.GetConfig().JWTSecret)

func GenerateJWT(email string, userId string) (tokenString string, err error) {
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &models.JWTClaim{
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

func ValidateToken(signedToken string) (claims *models.JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*models.JWTClaim)
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
