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

type RegisterVisitorService struct {
}

func (RegisterVisitorService) Execute(ctx context.Context, req *model.RegisterVisitorRequest) (resp *model.RegisterVisitorResponse, err error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 365).Unix()
	userID := cast.ToInt64(req.UserID)
	if userID == 0 {
		userID = config.GenerateIDInt64()
	}
	token, err := middleware.JwtDefaultClient.GenerateToken(val_obj.UserClaims{
		UserID:   userID,
		Nickname: "游客" + req.FingerPrint,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	})
	return &model.RegisterVisitorResponse{
		Meta: common.MetaOk,
		Data: model.RegisterVisitorData{
			Token:          token,
			TokenExpiresAt: expiresAt,
		},
	}, nil
}
