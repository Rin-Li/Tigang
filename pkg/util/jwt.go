package util

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var secret = []byte("secret-tigang")

func GenerateToken(UserID string)(string, error){
	claims := Claims{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate((time.Now().Add(time.Hour * 24))),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "Tigang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenString string)(*Claims, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token)(interface{}, error){
		return secret, nil
	})

	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid{
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

