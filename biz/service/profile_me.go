package service

import (
	"context"
	"nearby/biz/common"
	"nearby/biz/model"
)

type ProfileMeService struct {
}

func (ss *ProfileMeService) Execute(ctx context.Context, req *model.ProfileMeRequest) (resp *model.ProfileMeResponse, err error) {
	user := common.GetUser(ctx)
	resp = &model.ProfileMeResponse{
		Meta: common.MetaOk,
		User: model.User{
			UserID:   user.UserID,
			Nickname: user.Nickname,
			HeadURL:  user.HeadURL,
		},
	}
	return resp, nil
}
