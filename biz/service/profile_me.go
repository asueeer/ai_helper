package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"context"
	"github.com/spf13/cast"
)

type ProfileMeService struct {
}

func (ss *ProfileMeService) Execute(ctx context.Context, req *model.ProfileMeRequest) (resp *model.ProfileMeResponse, err error) {
	user := common.GetUser(ctx)
	resp = &model.ProfileMeResponse{
		Meta: common.MetaOk,
		Data: model.ProfileMeData{
			User: model.User{
				UserID:   cast.ToString(user.UserID),
				Nickname: user.Nickname,
				HeadURL:  user.HeadURL,
				IsHelper: user.IsHelper,
			},
		},
	}
	return resp, nil
}
