package val_obj

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	UserID   int64
	Nickname string
	HeadURL  string
	jwt.StandardClaims
}
