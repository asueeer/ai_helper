package entity

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/po"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/model/vo"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"time"
)

const (
	MsgHasNotRead = 1
	MsgHasRead    = 2
)

type MsgContent struct {
	Text     *string  `json:"text,omitempty"`
	RichText *string  `json:"rich_text,omitempty"`
	ImgURL   *string  `json:"img_url,omitempty"`
	AudioURL *string  `json:"audio_url,omitempty"`
	VideoURL *string  `json:"video_url,omitempty"`
	Link     *string  `json:"link,omitempty"`
	End      *bool    `json:"end,omitempty"`
	Options  []string `json:"options,omitempty"`
}

func (content *MsgContent) ToVo() vo.MsgContent {
	return vo.MsgContent{
		Text:     content.Text,
		RichText: content.RichText,
		ImgURL:   content.ImgURL,
		AudioURL: content.AudioURL,
		VideoURL: content.VideoURL,
		Options:  content.Options,
	}
}

type MessageFrom struct {
	ID         int64      `json:"id"`                    // 自增id, 无业务意义
	MessageID  int64      `json:"message_id"`            // 消息id
	ConvID     int64      `json:"conv_id"`               // 会话id
	SenderID   int64      `json:"sender_id"`             // 发送方id
	ReceiverID int64      `json:"receiver_id,omitempty"` // 接收方id
	Status     string     `json:"status,omitempty"`      // 消息状态
	Content    MsgContent `json:"content"`               // 消息内容
	Type       string     `json:"type"`                  // 消息类型
	SeqID      int64      `json:"seq_id"`                // 用于保序的序列号
	CreateAt   time.Time  `json:"create_at"`             // 创建时间戳
	Role       string     `json:"role"`                  // 发消息人的角色
}

func (f *MessageFrom) Persist(ctx context.Context) error {
	msgRepo := repo.NewMessageRepo()
	msgPo, err := f.ToPo()
	if err != nil {
		return err
	}
	msgFromPo, err := msgRepo.CreateMessageFrom(ctx, *msgPo)
	if err != nil {
		return err
	}
	f.CreateAt = msgFromPo.CreatedAt
	return nil
}

func NewMessageFromByPo(msgPo *po.MessageFrom) (*MessageFrom, error) {
	var content MsgContent
	err := json.Unmarshal(msgPo.Content, &content)
	if err != nil {
		return nil, err
	}
	return &MessageFrom{
		ID:         msgPo.ID,
		MessageID:  msgPo.MessageID,
		ConvID:     msgPo.ConvID,
		SenderID:   msgPo.SenderID,
		ReceiverID: msgPo.ReceiverID,
		Content:    content,
		Type:       msgPo.Type,
		SeqID:      msgPo.SeqID,
		CreateAt:   msgPo.CreatedAt,
	}, nil
}

func (f *MessageFrom) ToPo() (*po.MessageFrom, error) {
	j, err := json.Marshal(f.Content)
	if err != nil {
		return nil, err
	}
	return &po.MessageFrom{
		ID:         f.ID,
		MessageID:  f.MessageID,
		SenderID:   f.SenderID,
		ReceiverID: f.ReceiverID,
		Content:    j,
		Type:       f.Type,
		ConvID:     f.ConvID,
		SeqID:      f.SeqID,
	}, nil
}

func (f *MessageFrom) ToVo() *vo.Message {
	msgVo := &vo.Message{
		MessageID:  cast.ToString(f.MessageID),
		ConvID:     cast.ToString(f.ConvID),
		SenderID:   cast.ToString(f.SenderID),
		ReceiverID: cast.ToString(f.ReceiverID),
		Content: vo.MsgContent{
			Text: f.Content.Text,
			Link: f.Content.Link,
			End:  f.Content.End,
		},
		Type:   f.Type,
		Status: f.Status,
		SeqID:  f.SeqID,
	}
	if msgVo.SenderID == cast.ToString(common.HelperID) {
		msgVo.Role = common.ConvRoleHelper
	} else {
		msgVo.Role = common.ConvRoleVisitor
	}
	return msgVo
}

type MessageTo struct {
	ID        int64 `json:"id"`         // 自增id, 无业务意义
	MessageID int64 `json:"message_id"` // 消息id
	ConvID    int64 `json:"conv_id"`    // 会话id
	OwnerID   int64 `json:"owner_id"`   // 收件箱所有者id
	SeqID     int64 `json:"timestamp"`  // 用于保序的序列号
}

func NewMessageToByPo(po *po.MessageTo) *MessageTo {
	return &MessageTo{
		ID:        po.ID,
		MessageID: po.MessageID,
		ConvID:    po.ConvID,
		OwnerID:   po.OwnerID,
		SeqID:     po.SeqID,
	}
}

func (t *MessageTo) Persist(ctx context.Context) error {
	msgRepo := repo.NewMessageRepo()
	_, err := msgRepo.CreateMessageTo(ctx, t.ToPo())
	if err != nil {
		return err
	}
	return nil
}

func (t *MessageTo) ToPo() po.MessageTo {
	return po.MessageTo{
		ID:        t.ID,
		MessageID: t.MessageID,
		ConvID:    t.ConvID,
		OwnerID:   t.OwnerID,
		SeqID:     t.SeqID,
	}
}
