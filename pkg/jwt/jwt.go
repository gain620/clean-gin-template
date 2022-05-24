package jwt

import (
	"clean-gin-template/pkg/logger"
	"fmt"
	"gopkg.in/square/go-jose.v2/jwt"
	"strings"
)

type JwtValidator interface {
	Validate(tokenStr string) (*jwt.Token, error)
}

type jwtValidator struct {
	l logger.Interface
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

func New(l logger.Interface) JwtValidator {
	return &jwtValidator{
		l: l,
	}
}

func (j jwtValidator) Validate(tokenStr string) (*jwt.Token, error) {
	//return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	//
	//	// Signing method validation
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//	}
	//
	//	// Return the secret signing key
	//	//return []byte(jwtSrv.secretKey), nil
	//})
}

func (j jwtValidator) CheckJwtRole(tokenStr, roleName string) (bool, error) {
	var claims JwtPayload // custom payload struct to store parsed token

	// decode JWT token without verifying the signature
	token, _ := jwt.ParseSigned(tokenStr)
	_ = token.UnsafeClaimsWithoutVerification(&claims)

	for _, item := range claims.Roles {
		_, jwtRole, ok := strings.Cut(item, "_")
		if ok != true {
			j.l.Fatal("the format does not fit")
			return false, fmt.Errorf("the role format does not fit")
		}

		if jwtRole == roleName {
			j.l.Print("ACCESS ALLOWED ROLE")
			return true, nil
		}
	}

	return false, nil
}
