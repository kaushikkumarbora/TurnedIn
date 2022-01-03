package model

import "github.com/golang-jwt/jwt"

type Auth struct {
	Id       int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
