package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// JwtPayLoad jwt 负载
type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 权限  1 普通用户  2 管理员
}

// CustomClaims 自定义声明
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

func main() {
	// 生成token
	//CreatFullToken(
	//	JwtPayLoad{
	//		UserID:   1,
	//		Username: "admin",
	//		Role:     2,
	//	}, "123", 100)
	//
	//CreatOptionalToken(JwtPayLoad{
	//	UserID:   1,
	//	Username: "admin",
	//	Role:     2,
	//}, "123", 100)

	ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6MiwiaXNzIjoidGVzdCIsImV4cCI6MTczNjMyNTcwMH0.wpNkZdzB1vwxY09tY20_6eYuCYY-Hq6_rONfnVl6ruA", "123")
}

func CreatFullToken(user JwtPayLoad, accessSecret string, expires int64) {
	// Create the Claims
	// 完整的声明
	claims := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			Issuer:    "test",                                             // 签发人
			Subject:   "somebody",                                         // 主题
			ID:        "1",                                                // 编号
			Audience:  []string{"somebody_else"},                          // 受众
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用HS256算法
	ss, err := token.SignedString([]byte(accessSecret))        //  加盐签名
	fmt.Printf("%v %v", ss, err)
}

func CreatOptionalToken(user JwtPayLoad, accessSecret string, expires int64) {
	claims := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			Issuer:    "test",                                             // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用HS256算法
	ss, err := token.SignedString([]byte(accessSecret))        // 签名
	fmt.Printf("%v %v", ss, err)
}

func ParseToken(tokenStr string, accessSecret string) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}

func ParseTokenCustomValidation(tokenStr string, accessSecret string) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}
