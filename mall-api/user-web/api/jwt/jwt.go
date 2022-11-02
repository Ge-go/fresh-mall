package jwttk

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"mall-api/user-web/global"
	jwtmodel "mall-api/user-web/global/jwt"
	"time"
)

var (
	TokenInvalid = errors.New("couldn't handle this token")
)

type JWTToken struct {
	SigningKey *rsa.PrivateKey
}

func NewJWTToken() *JWTToken {
	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(global.ServerConfig.JWTInfo.SigningKey))
	return &JWTToken{
		privateKey,
	}
}

// 服务拆分,这里只进行jwt token加密

// CreateToken 创建一个token
func (j *JWTToken) CreateToken(claims jwtmodel.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	return token.SignedString(j.SigningKey)
}

// RefreshToken 更新token
func (j *JWTToken) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwtmodel.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*jwtmodel.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
