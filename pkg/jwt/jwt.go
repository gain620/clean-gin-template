package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JwtValidator interface {
	Validate(tokenStr string) (*jwt.Token, error)
}

type jwtValidator struct {
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

func New() JwtValidator {
	return &jwtValidator{}
}

func (j jwtValidator) Validate(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}
