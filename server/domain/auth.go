package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"userId"`
}
