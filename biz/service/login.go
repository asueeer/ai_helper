package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/config"
	"ai_helper/biz/domain/val_obj"
	"ai_helper/biz/middleware"
	"ai_helper/biz/model"
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"time"
)

type LoginService struct {
}

var record = map[string]*val_obj.UserClaims{
	"aaa": {
		UserID:   9999999,
		Nickname: "aaa",
		HeadURL:  "",
	},
	"bbb": {
		UserID:   9999998,
		Nickname: "bbb",
		HeadURL:  "",
	},
	"ccc": {
		UserID:   9999997,
		Nickname: "ccc",
	},
}

func (LoginService) Execute(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	if req.Password != "123" || record[req.Username] == nil {
		return nil, common.NewBizErr(common.BizErrCode, "wrong pwd or username", nil)
	}
	user := record[req.Username]
	expiresAt := time.Now().Add(time.Hour * 24 * 365).Unix()
	userID := cast.ToInt64(user)
	if userID == 0 {
		userID = config.GenerateIDInt64()
	}
	token, _ := middleware.JwtDefaultClient.GenerateToken(val_obj.UserClaims{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	})
	return &model.LoginResponse{
		Meta: common.MetaOk,
		Data: model.LoginData{
			Token:          token,
			TokenExpiresAt: expiresAt,
		},
	}, nil
}
