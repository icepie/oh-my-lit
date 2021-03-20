package util

import (
	"time"

	"github.com/icepie/lit-edu-go/conf"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(conf.ProConf.JWT.Secret)

//Claims ...
type Claims struct {
	Info string `json:"info"`
	jwt.StandardClaims
}

//GenerateToken 签发Token
func GenerateToken(Info string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(conf.ProConf.JWT.ExpireTime) * time.Hour)

	claims := Claims{
		Info,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "singzer",
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
