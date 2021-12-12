package entity

import (
	"ai_helper/biz/model/vo"
	"ai_helper/biz/util"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"time"

	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/po"
	"ai_helper/biz/dal/db/repo"
)

type Participant struct {
	UserID   int64  `json:"user_id"`
	HeadURL  string `json:"head_url"`
	Nickname string `json:"nickname"`
}

type Conversation struct {
	ID           int64        `json:"id"`               // 主键自增id,无业务意义
	ConvID       int64        `json:"conv_id"`          // 会话id, 会话的唯一标识
	Type         string       `json:"type"`             // 会话类型
	Creator      int64        `json:"creator"`          // 创建者id
	Acceptor     int64        `json:"acceptor"`         // 接收者id
	Status       string       `json:"status,omitempty"` // 状态
	LastMsgID    int64        `json:"last_msg_id"`      // 最近一条消息的msg_id
	LastMsg      *MessageFrom `json:"last_msg"`         // 最新的一条消息
	Participants []int64      `json:"participants"`     // 会话参与者
	Timestamp    time.Time    `json:"timestamp"`        // 会话时间戳
	SeqID        int64        `json:"seq_id"`           // 序列号, 用于保序
	UnreadCnt    int          `json:"unread_cnt"`       // 未读消息数
}

func (c *Conversation) ToPo() po.Conversation {
	return po.Conversation{
		ID:        c.ID,
		ConvID:    c.ConvID,
		Type:      c.Type,
		Creator:   c.Creator,
		Status:    c.Status,
		Timestamp: c.Timestamp,
		SeqID:     util.Sec2Mirco(c.Timestamp.Unix()),
	}
}

func (c *Conversation) IsHelperType() bool {
	return c.Type == common.HelperConversationType
}

func GetConversationEntityByID(ctx context.Context, convID int64) (*Conversation, error) {
	convRepo := repo.NewConversationRepo()
	convPo, err := convRepo.GetConvPo(ctx,
		repo.GetConvPoRequest{
			ConvID: convID,
		},
	)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return NewConversationEntityByPo(ctx, convPo), err
}

// NewConversationEntityByPo 根据po构造entity
func NewConversationEntityByPo(ctx context.Context, convPo *po.Conversation) *Conversation {
	if convPo == nil {
		return &Conversation{}
	}
	return &Conversation{
		ID:        convPo.ID,
		ConvID:    convPo.ConvID,
		Type:      convPo.Type,
		Creator:   convPo.Creator,
		Status:    convPo.Status,
		Timestamp: convPo.Timestamp,
		LastMsgID: convPo.LastMsgID,
		Acceptor:  convPo.Acceptor,
	}
}

type GetMessagesRequest struct {
	Limit     int64 `json:"limit"`
	SeqIDFrom int64 `json:"seq_id_from"`
	SeqIDTo   int64 `json:"seq_id_to"`
	ViewerID  int64 `json:"viewer_id"`
}

type GetMessagesResponse struct {
	Total    int64             `json:"total"`
	MsgTos   []*po.MessageTo   `json:"msg_tos"`
	MsgFroms []*po.MessageFrom `json:"msg_froms"`
}

func (c *Conversation) GetMessages(ctx context.Context, req GetMessagesRequest) (*GetMessagesResponse, error) {
	msgRepo := repo.NewMessageRepo()
	resp := &GetMessagesResponse{}
	// 查询收件箱里的消息
	msgTos, total, err := msgRepo.GetMessageTos(ctx, repo.GetMessagesRequest{
		ConvID:    c.ConvID,
		Limit:     req.Limit,
		SeqIDFrom: req.SeqIDFrom,
		SeqIDTo:   req.SeqIDTo,
		OwnerID:   req.ViewerID,
	})
	if err != nil {
		return nil, err
	}
	// 查询发件箱里的消息，便于上层组装信息
	msgFroms, err := GetMessageFroms(ctx, msgTos)
	if err != nil {
		return nil, err
	}
	resp = &GetMessagesResponse{
		Total:    total,
		MsgFroms: msgFroms,
		MsgTos:   msgTos,
	}
	return resp, nil
}

func (c *Conversation) ToVo() *vo.Conversation {
	return &vo.Conversation{
		ConvID:    cast.ToString(c.ConvID),
		Type:      c.Type,
		UnRead:    cast.ToInt32(c.UnreadCnt),
		Timestamp: util.Sec2Mirco(c.Timestamp.Unix()),
		Status:    c.Status,
	}
}

func (c *Conversation) ReOpen(ctx context.Context) error {
	// 返回前更改会话单的状态
	convRepo := repo.NewConversationRepo()
	err := convRepo.UpdateConvStatus(ctx, repo.UpdateConvStatusRequest{
		ConvID:    c.ConvID,
		Status:    "waiting",
		PreStatus: "end",
	})
	return err
}

func GetMessageFroms(ctx context.Context, messageTos []*po.MessageTo) ([]*po.MessageFrom, error) {
	if len(messageTos) == 0 {
		return nil, nil
	}
	msgIDs := make([]int64, len(messageTos))
	for i := range msgIDs {
		msgIDs[i] = messageTos[i].MessageID
	}
	msgRepo := repo.NewMessageRepo()
	msgFroms, err := msgRepo.GetMessageFroms(ctx, repo.GetMessageFromsRequest{
		MessageIDs: msgIDs,
	})
	if err != nil {
		return nil, err
	}
	return msgFroms, nil
}
