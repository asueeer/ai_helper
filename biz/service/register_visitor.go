package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"nearby/biz/common"
	"nearby/biz/domain/val_obj"
	"nearby/biz/middleware"
	"nearby/biz/model"
	"time"
)

type RegisterVisitorService struct {
}

func (RegisterVisitorService) Execute(ctx context.Context, req *model.RegisterVisitorRequest) (resp *model.RegisterVisitorResponse, err error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 365).Unix()
	token, err := middleware.JwtDefaultClient.GenerateToken(val_obj.UserClaims{
		UserID:   cast.ToInt64(req.UserID),
		Nickname: "游客" + req.UserID,
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
