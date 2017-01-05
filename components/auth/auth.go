package auth

import (
	//"gopkg.in/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var tokenEncodeString string = "something"

func GenTokenV2() (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return (token.SignedString(mySigningKey))
}

func ParseToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
