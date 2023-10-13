package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/wonderivan/logger"
	"time"
)

var JWTToken jwtToken

type jwtToken struct{}

// CustomClaims 定义token 携带的信息
type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	*jwt.StandardClaims
}

// SECRET 加解密因子
const (
	SECRET = "devops"
)

// ParseToken 解析token
func (*jwtToken) ParseToken(tokenString string) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		logger.Error("Token parse failed" + err.Error())
		//处理 token 的各种报错
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("TokenMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("TokenExpired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("TokenNotValidYet")
			} else {
				return nil, errors.New("TokenInvalid")
			}
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("解析Token失败")
}

// CreateJwtToken 后端产生token 用于验证jwt 中间件是否生效
func CreateJwtToken(username string, password string) (string, error) {

	// 定义过期时间,7天后过期
	expireToken := time.Now().Add(time.Hour * 24 * 7).Unix()

	claims := &CustomClaims{
		username,
		password,
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // token信息生效时间
			ExpiresAt: int64(expireToken),              // 过期时间
			Issuer:    "wonders",                       // 发布者
		},
	}
	// 对自定义claims加密,jwt.SigningMethodHS256是加密算法得到第二部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 给这个token盐加密 第三部分,得到一个完整的三段的加密
	signedToken, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
