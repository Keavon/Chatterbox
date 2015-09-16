package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token type for jwt tokens.
type Token struct {
	Secret string
	ISS    string
}

// New creates a new jwt token.
func (t Token) New(id string, exp time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["sub"] = id
	token.Claims["iss"] = t.ISS
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = exp.Unix()
	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(t.Secret))
}

// Valid checks if a token is valid
func (t Token) Valid(rawToken string) (bool, string) {
	parsedToken, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return token, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(t.Secret), nil
	})

	if err == nil && parsedToken.Valid {
		return true, parsedToken.Claims["sub"].(string)
	}

	return false, ""
}
