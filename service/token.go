package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"suning/model"
	"time"
)

var Secret = []byte("YJX")

// keyFunc 用于验证签名是否正确
func keyFunc(*jwt.Token) (i interface{}, err error) {
	return Secret, nil
}

// GenToken GenToken生成aToken和rToken
func GenToken(uid, role string) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := model.MyClaims{
		UserId: uid,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // 过期时间
			Issuer:    "YJX",                            // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(Secret)
	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间
		Issuer:    "YJX",                                 // 签发人
	}).SignedString(Secret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// RefreshToken 刷新aToken
func RefreshToken(rToken string) (newAToken, newRToken string, err error) {
	var claims model.MyClaims
	_, err = jwt.ParseWithClaims(rToken, &claims, keyFunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", "", errors.New("invalid refresh token signature")
		}
	}
	//生成新的token
	newAToken, newRToken, err = GenToken(claims.UserId, claims.Role)
	if err != nil {
		return "", "", err
	}
	return newAToken, newRToken, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
