package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"context"
	"github.com/spf13/cast"
)

type SendMessageToRobotService struct {
}

func (SendMessageToRobotService) Execute(ctx context.Context, req *model.SendMessageRequest) (*model.SendMessageToRobotResponse, error) {
	return &model.SendMessageToRobotResponse{
		SendMessageData: model.SendMessageToRobotData{
			RespContent: cast.ToString(req.Content.Text),
		},
		Meta: common.MetaOk,
	}, nil
}
