package jwt

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
