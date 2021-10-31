package entity

import (
	"nearby/biz/dal/db/po"
	"nearby/biz/domain/val_obj"
	"nearby/biz/middleware"
	"nearby/biz/model/vo"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       int64  `json:"id"`       // 主键自增id,无业务意义
	UserID   int64  `json:"user_id"`  // 用户id, 用户唯一标示
	Nickname string `json:"nickname"` // 用户昵称
	HeadURL  string `json:"head_url"` // 用户头像链接
}

func (user *User) GenerateToken() (*val_obj.Token, error) {
	cli := middleware.JwtDefaultClient
	expiresAt := time.Now().Add(time.Hour * 24 * 90).Unix()
	userClaims := val_obj.UserClaims{
		Nickname: user.Nickname,
		HeadURL:  user.HeadURL,
		UserID:   user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    cli.Issuer,
			Subject:   cli.Subject,
		},
	}
	token, err := cli.GenerateToken(userClaims)
	if err != nil {
		return nil, err
	}
	return &val_obj.Token{
		Token:          token,
		TokenExpiredAt: expiresAt,
	}, nil
}

func NewUserEntityByPo(po *po.User) *User {
	if po == nil {
		return nil
	}
	userEntity := &User{
		ID:       po.ID,
		UserID:   po.UserID,
		Nickname: po.Nickname,
		HeadURL:  po.HeadURL,
	}
	return userEntity
}

func (user *User) ToVO() *vo.User {
	return &vo.User{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		HeadURL:  user.HeadURL,
	}
}

func (user *User) ToPo() *po.User {
	return &po.User{
		ID:       user.ID,
		UserID:   user.UserID,
		Nickname: user.Nickname,
		HeadURL:  user.HeadURL,
	}
}
