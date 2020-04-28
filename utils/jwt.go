package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID string, roleID string, jwtSigned string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[JWT_USER_ID] = userID
	claims[JWT_ROLE_ID] = roleID
	claims[JWT_EXPIRED] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte(jwtSigned))

	return tokenString
}

func IsValidToken(tokenString string, jwtSigned string) (bool, string, string, error) {
	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSigned), nil
	}); token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return true, claims[JWT_USER_ID].(string), claims[JWT_ROLE_ID].(string), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		var errMessage error
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errMessage = fmt.Errorf("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errMessage = fmt.Errorf("Token Expired")
		} else {
			errMessage = fmt.Errorf("Couldn't handle this token: %q", err)
		}
		return false, "", "", errMessage
	} else {
		errMessage := fmt.Errorf("Couldn't handle this token: %q", err)
		return false, "", "", errMessage
	}
}
