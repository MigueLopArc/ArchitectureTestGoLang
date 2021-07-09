package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
