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

func (LoginService) Execute(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	if req.Password != "123" || common.Record[req.Username] == nil {
		return nil, common.NewBizErr(common.BizErrCode, "wrong pwd or username", nil)
	}
	user := common.Record[req.Username]
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
