package po

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

// MessageFrom 消息发件箱, 存储发送者角度的信息
type MessageFrom struct {
	ID         int64           `gorm:"column:id"`          // 主键自增id,无业务意义
	MessageID  int64           `gorm:"column:message_id"`  // 消息id
	ConvID     int64           `gorm:"column:conv_id"`     // 所属会话id
	SenderID   int64           `gorm:"column:sender_id"`   // 发送者的id
	ReceiverID int64           `gorm:"column:receiver_id"` // 接收者的id
	Content    json.RawMessage `gorm:"column:content"`     // 消息内容
	Type       string          `gorm:"column:type"`        // 消息类型
	Timestamp  time.Time       `gorm:"column:timestamp"`   // 消息时间戳

	gorm.Model
}

func (MessageFrom) TableName() string {
	return "message_from"
}

// 消息收件箱，存储收件者角度的信息，如已读/未读等。
type MessageTo struct {
	ID        int64     `gorm:"column:id"`         // 主键自增id,无业务意义
	MessageID int64     `gorm:"column:message_id"` // 消息id
	ConvID    int64     `gorm:"column:conv_id"`    // 会话id
	OwnerID   int64     `gorm:"column:owner_id"`   // 收件箱所有者id
	TimeStamp time.Time `gorm:"column:timestamp"`  // 消息时间戳
	HasRead   int32     `gorm:"column:has_read"`   //  消息是否已读; 1: 未读; 2: 已读

	gorm.Model
}

func (MessageTo) TableName() string {
	return "message_to"
}
