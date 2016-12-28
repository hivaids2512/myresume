package auth

import (
	"../../modules/user"
	"gopkg.in/dgrijalva/jwt-go.v2"
)

var tokenEncodeString string = "something"

func GenToken(user User) (token, err) {
	token := jwt.New(jwt.SigningMethodHS256)

	// set some claims
	token.Claims["email"] = user.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//Sign and get the complete encoded token as string
	return (token.SignedString(tokenEncodeString))
}

func ParseToken(myToken string) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return myLookupKey(token.Header["kid"]), nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
