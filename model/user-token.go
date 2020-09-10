package model

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	LogginRequestID string
	UserID int
	UserAgent string
	Source string
	LoginTime int
	jwt.StandardClaims
}