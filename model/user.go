package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Uid      int    `json:"uid" form:"uid"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type MyClaims struct {
	Uid  string `json:"uid"`
	Role string `json:"role"`
	jwt.StandardClaims
}
