package middleware

import (
	"nearby/biz/config"

	"nearby/biz/domain/val_obj"

	"github.com/golang-jwt/jwt"
)

var JwtDefaultClient *JwtClient

func init() {
	jwtConf := config.JwtDefaultConfig
	JwtDefaultClient = NewJwtClient(jwtConf.JwtKey, jwtConf.Issuer, jwtConf.Subject)
}

func NewJwtClient(jwtKey string, issuer string, subject string) *JwtClient {
	return &JwtClient{
		JwtKey:  jwtKey,
		Issuer:  issuer,
		Subject: subject,
	}
}

type JwtClient struct {
	JwtKey  string
	Issuer  string
	Subject string
}

// ParseToken 解析token
func (c *JwtClient) ParseToken(tokenString string) (*jwt.Token, *val_obj.UserClaims, error) {
	claims := &val_obj.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(c.JwtKey), nil
	})
	return token, claims, err
}

func (c *JwtClient) GenerateToken(claims val_obj.UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.JwtKey))
}
