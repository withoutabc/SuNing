package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username" `
	Password string `json:"password" `
}

type MyClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
