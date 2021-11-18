package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"context"
	"encoding/json"
)

type WsService struct {
}

func (ss *WsService) Wrap(err error) []byte {
	bizErr := common.NewBizErr(common.BizErrCode, "解析json出错了", err)
	return []byte(bizErr.Error())
}

func (ss *WsService) WsLoadConv(ctx context.Context, msgType int, msg []byte) []byte {
	convService := LoadConversationDetailService{}
	var req *model.LoadConversationDetailRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return ss.Wrap(err)
	}
	resp, err := convService.Execute(ctx, req)
	if err != nil {
		return ss.Wrap(err)
	}
	wsResp := model.WsMessageResponse{
		Type: msgType,
		Msg:  resp,
	}
	j, err := json.Marshal(wsResp)
	if err != nil {
		return ss.Wrap(err)
	}
	return j
}

func (ss *WsService) WsLoadConvs(ctx context.Context, msgType int, msg []byte) []byte {
	convService := LoadConversationsService{}
	var req *model.LoadConversationsRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return ss.Wrap(err)
	}
	resp, err := convService.Execute(ctx, req)
	if err != nil {
		return ss.Wrap(err)
	}
	wsResp := model.WsMessageResponse{
		Type: msgType,
		Msg:  resp,
	}
	j, err := json.Marshal(wsResp)
	if err != nil {
		return ss.Wrap(err)
	}
	return j
}
