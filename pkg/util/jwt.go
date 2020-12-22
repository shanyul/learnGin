package util

import (
	"fmt"
	"learngo/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// jwtSecret 密匙
var jwtSecret = []byte(setting.JwtSecret)

// Claims 登录信息
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(username, password string) (string, error) {
	fmt.Println(jwtSecret)
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims {
            ExpiresAt : expireTime.Unix(),
            Issuer : "learngo",
        },
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}