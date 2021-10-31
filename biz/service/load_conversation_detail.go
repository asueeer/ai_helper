package service

import (
	"context"
	"math"
	"nearby/biz/common"
	domainService "nearby/biz/domain/service"
	"nearby/biz/model"
	"nearby/biz/model/vo"
	"nearby/biz/util"
	"time"

	"github.com/spf13/cast"
)

type LoadConversationDetailService struct {
}

func (ss *LoadConversationDetailService) checkParams(ctx context.Context, req *model.LoadConversationDetailRequest) error {
	return nil
}

func (ss *LoadConversationDetailService) checkAuth(ctx context.Context, convID int64) error {
	return nil
}

func (ss *LoadConversationDetailService) Execute(ctx context.Context, req *model.LoadConversationDetailRequest) (resp *model.LoadConversationDetailResponse, err error) {
	if req.Limit == 0 {
		req.Limit = 40
	}
	if err = ss.checkParams(ctx, req); err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "参数校验错误", err)
	}
	// 校验当前用户有没有权限查看当前会话
	if ss.checkAuth(ctx, req.ConvID) != nil {
		return nil, common.NewBizErr(common.EvilViewErrCode, "ops, 会话找不到了...", err)
	}
	var conversationLoader domainService.ConversationLoadService
	if req.Cursor == 0 {
		req.Cursor = time.Now().Unix()
	}
	timestampTo := cast.ToTime(req.Cursor)

	viewerID, err := ss.GetViewerID(ctx, req)
	if err != nil {
		return nil, common.NewBizErr(common.EvilViewErrCode, "ops, 会话找不到了...", err)
	}
	convResp, err := conversationLoader.GetConversation(ctx, domainService.GetConversationRequest{
		ConvID:      req.ConvID,
		Limit:       req.Limit,
		TimestampTo: &timestampTo,
		ViewerID:    viewerID,
	})
	if err != nil {
		return nil, err
	}
	msgVos, newCursor, err := ss.ConstructMsgVos(ctx, convResp)
	if err != nil {
		return nil, err
	}
	resp = &model.LoadConversationDetailResponse{
		Meta: common.MetaOk,
		Data: model.LoadConversationDetailData{
			Messages:  msgVos,
			NewCursor: newCursor,
		},
	}
	if len(resp.Data.Messages) == cast.ToInt(req.Limit) {
		resp.Data.HasMore = true
	}
	if resp.Data.NewCursor == req.Cursor {
		resp.Data.HasMore = false
	}
	if !resp.Data.HasMore {
		resp.Data.NewCursor = 0
	}
	return resp, nil
}

func (ss *LoadConversationDetailService) ConstructMsgVos(ctx context.Context, resp *domainService.GetConversationResponse) (vos []*vo.Message, newCursor int64, err error) {
	msgAggs := resp.MsgAggs
	vos = make([]*vo.Message, len(msgAggs))
	// 设置哨兵浮标
	// 若浮标没有改变, 说明再滚动拉取也不会有新消息了
	newCursor = math.MaxInt64
	for i := range vos {
		msgFrom := msgAggs[i].MessageFrom
		vos[i] = &vo.Message{
			SenderID:   msgFrom.SenderID,
			ReceiverID: msgFrom.ReceiverID,
			Content:    msgFrom.Content.ToVo(),
			Type:       msgFrom.Type,
			Status:     msgFrom.Status,
			Timestamp:  msgFrom.Timestamp.Unix(),
		}
		newCursor = util.Min64(newCursor, msgFrom.Timestamp.Unix())
	}

	return vos, newCursor, nil
}

func (ss *LoadConversationDetailService) GetViewerID(ctx context.Context, req *model.LoadConversationDetailRequest) (viewerID int64, err error) {
	if req.Role == common.HelperConversationType {
		// fixme 没有做权限校验
		return common.HelperID, nil
	}
	user := common.GetUser(ctx)
	return user.UserID, nil
}
