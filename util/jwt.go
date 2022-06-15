package util

import (
	"time"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Info interface{} `json:"info"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(generateMd5(global.JWT.Secret))
}

// 生成token
func GenerateToken(info interface{}) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWT.Expire)
	claims := Claims{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWT.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
