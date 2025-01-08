package utils

import "github.com/golang-jwt/jwt/v5"

func GenerateToken(data interface{}) {
	jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		
	})
}
