package service

import (
	"ai_helper/biz/common"
	domainService "ai_helper/biz/domain/service"
	"ai_helper/biz/model"
	"ai_helper/biz/model/vo"
	"ai_helper/biz/util"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"math"
	"time"
)

type LoadConversationDetailService struct {
}

func (ss *LoadConversationDetailService) checkParams(ctx context.Context, req *model.LoadConversationDetailRequest) error {
	if req.Limit == 0 {
		req.Limit = 50
	}
	invalidCursor := req.Cursor == "0" || req.Cursor == ""
	if invalidCursor {
		if req.Direction == -1 {
			req.Cursor = cast.ToString(util.Sec2Mirco(time.Now().Unix()))
		}
		if req.Direction == 1 {
			return errors.New("需要一个明确的同步位点")
		}
	}
	if req.Direction == 0 {
		req.Direction = -1
	}
	return nil
}

func (ss *LoadConversationDetailService) checkAuth(ctx context.Context, convID int64) error {
	return nil
}

func (ss *LoadConversationDetailService) Execute(ctx context.Context, req *model.LoadConversationDetailRequest) (resp *model.LoadConversationDetailResponse, err error) {
	if err = ss.checkParams(ctx, req); err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "参数校验错误", err)
	}
	// 校验当前用户有没有权限查看当前会话
	if ss.checkAuth(ctx, cast.ToInt64(req.ConvID)) != nil {
		return nil, common.NewBizErr(common.EvilViewErrCode, "ops, 会话找不到了...", err)
	}
	var conversationLoader domainService.ConversationLoader

	var seqIDFrom, seqIDTo int64
	if req.Direction == 1 {
		seqIDFrom = cast.ToInt64(req.Cursor)
	}
	if req.Direction == -1 {
		seqIDTo = cast.ToInt64(req.Cursor)
	}

	viewerID, err := ss.GetViewerID(ctx, req)
	if err != nil {
		return nil, common.NewBizErr(common.EvilViewErrCode, "ops, 会话找不到了...", err)
	}
	convResp, err := conversationLoader.GetConversation(ctx, domainService.GetConversationRequest{
		ConvID:    cast.ToInt64(req.ConvID),
		Limit:     cast.ToInt64(req.Limit),
		SeqIDTo:   seqIDTo,
		SeqIDFrom: seqIDFrom,
		ViewerID:  viewerID,
	})
	if err != nil {
		return nil, err
	}
	msgVos, newCursor, err := ss.ConstructMsgVos(ctx, convResp, req.Direction)
	// 很恶心的业务逻辑, 等着要解耦
	if newCursor == -1 {
		newCursor = cast.ToInt64(req.Cursor)
	}

	if err != nil {
		return nil, err
	}
	resp = &model.LoadConversationDetailResponse{
		Meta: common.MetaOk,
		Data: model.LoadConversationDetailData{
			Messages:  msgVos,
			NewCursor: cast.ToString(newCursor),
		},
	}
	if len(resp.Data.Messages) < cast.ToInt(req.Limit) {
		resp.Data.HasMore = false
	}
	if resp.Data.NewCursor == req.Cursor {
		resp.Data.HasMore = false
	}
	return resp, nil
}

func (ss *LoadConversationDetailService) ConstructMsgVos(ctx context.Context,
	resp *domainService.GetConversationResponse, direction int32) (
	vos []*vo.Message, newCursor int64, err error) {
	msgAggs := resp.MsgAggs
	vos = make([]*vo.Message, len(msgAggs))
	// 设置哨兵浮标
	// 若浮标没有改变, 说明再滚动拉取也不会有新消息了
	{
		newCursor = math.MaxInt64
		if direction == 1 {
			newCursor = -1
		}
	}

	for i := range vos {
		msgFrom := msgAggs[i].MessageFrom
		vos[i] = &vo.Message{
			MessageID:  cast.ToString(msgFrom.MessageID),
			SenderID:   cast.ToString(msgFrom.SenderID),
			ReceiverID: cast.ToString(msgFrom.ReceiverID),
			Content:    msgFrom.Content.ToVo(),
			Type:       msgFrom.Type,
			Status:     msgFrom.Status,
			SeqID:      msgFrom.SeqID,
		}
		if msgFrom.SenderID == common.HelperID {
			vos[i].Role = common.ConvRoleHelper
		} else {
			vos[i].Role = common.ConvRoleVisitor
		}
		{
			if direction == -1 {
				newCursor = util.MinInt64(newCursor, msgFrom.SeqID)
			}
			if direction == 1 {
				newCursor = util.MaxInt64(newCursor, msgFrom.SeqID)
			}
		}

	}
	// 应前端的要求，展示的时候是从上至下展示，所以要reverse一下顺序
	vos = ss.Reverse(vos)
	return vos, newCursor, nil
}

func (ss *LoadConversationDetailService) Reverse(vos []*vo.Message) []*vo.Message {
	ret := make([]*vo.Message, len(vos))
	for i := 0; i < len(vos); i++ {
		ret[i] = vos[len(vos)-i-1]
	}
	return ret
}

func (ss *LoadConversationDetailService) GetViewerID(ctx context.Context, req *model.LoadConversationDetailRequest) (viewerID int64, err error) {
	if req.Role == common.HelperConversationType {
		// fixme 没有做权限校验
		return common.HelperID, nil
	}
	user := common.GetUser(ctx)
	return user.UserID, nil
}
