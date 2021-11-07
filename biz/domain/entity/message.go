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
	Text     *string `json:"text,omitempty"`
	RichText *string `json:"rich_text,omitempty"`
	ImgURL   *string `json:"img_url,omitempty"`
	AudioURL *string `json:"audio_url,omitempty"`
	VideoURL *string `json:"video_url,omitempty"`
}

func (content *MsgContent) ToVo() vo.MsgContent {
	return vo.MsgContent{
		Text:     content.Text,
		RichText: content.RichText,
		ImgURL:   content.ImgURL,
		AudioURL: content.AudioURL,
		VideoURL: content.VideoURL,
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
	Timestamp  time.Time  `json:"timestamp"`             // 消息时间戳
}

func (f *MessageFrom) Persist(ctx context.Context) error {
	msgRepo := repo.NewMessageRepo()
	msgPo, err := f.ToPo()
	if err != nil {
		return err
	}
	_, err = msgRepo.CreateMessageFrom(ctx, *msgPo)
	if err != nil {
		return err
	}
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
		Timestamp:  msgPo.Timestamp,
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
		Timestamp:  f.Timestamp,
	}, nil
}

func (f *MessageFrom) ToVo() *vo.Message {
	msgVo := &vo.Message{
		MessageID:  cast.ToString(f.MessageID),
		SenderID:   cast.ToString(f.SenderID),
		ReceiverID: cast.ToString(f.ReceiverID),
		Content: vo.MsgContent{
			Text: f.Content.Text,
		},
		Type:      f.Type,
		Status:    f.Status,
		Timestamp: f.Timestamp.Unix() * 1000,
	}
	if msgVo.SenderID == cast.ToString(common.HelperID) {
		msgVo.Role = common.ConvRoleHelper
	} else {
		msgVo.Role = common.ConvRoleVisitor
	}
	return msgVo
}

type MessageTo struct {
	ID        int64     `json:"id"`         // 自增id, 无业务意义
	MessageID int64     `json:"message_id"` // 消息id
	ConvID    int64     `json:"conv_id"`    // 会话id
	OwnerID   int64     `json:"owner_id"`   // 收件箱所有者id
	Timestamp time.Time `json:"timestamp"`  // 消息时间戳
	HasRead   int32     `json:"has_read"`   // 是否已读, 1为未读; 2为已读
}

func NewMessageToByPo(po *po.MessageTo) *MessageTo {
	return &MessageTo{
		ID:        po.ID,
		MessageID: po.MessageID,
		ConvID:    po.ConvID,
		OwnerID:   po.OwnerID,
		Timestamp: po.TimeStamp,
		HasRead:   po.HasRead,
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
		TimeStamp: t.Timestamp,
		HasRead:   t.HasRead,
	}
}
