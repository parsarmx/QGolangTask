package models

import "github.com/golang-jwt/jwt"

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var JWTSecret = []byte("secret")
