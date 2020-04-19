package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(payload string, jwtSigned string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = payload
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte(jwtSigned))

	return tokenString
}

func IsValidToken(tokenString string, jwtSigned string) (bool, string, error) {
	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSigned), nil
	}); token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return true, claims["id"].(string), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		var errMessage error
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errMessage = fmt.Errorf("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errMessage = fmt.Errorf("Token Expired")
		} else {
			errMessage = fmt.Errorf("Couldn't handle this token: %q", err)
		}
		return false, "", errMessage
	} else {
		errMessage := fmt.Errorf("Couldn't handle this token: %q", err)
		return false, "", errMessage
	}
}
