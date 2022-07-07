package model

import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	ID int64 `json:"id"`
	jwt.StandardClaims
}
