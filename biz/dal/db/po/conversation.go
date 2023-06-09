package po

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Conversation 会话表
type Conversation struct {
	ID        int64     `gorm:"column:id"`          // 主键自增id,无业务意义
	ConvID    int64     `gorm:"column:conv_id"`     // 会话id
	Type      string    `gorm:"column:type"`        // 会话类型
	Creator   int64     `gorm:"column:creator"`     // 创建人id
	Acceptor  int64     `gorm:"column:acceptor"`    // 接收者id
	Status    string    `gorm:"column:status"`      // 会话状态
	LastMsgID int64     `gorm:"column:last_msg_id"` // 最近一条消息的msg_id
	Timestamp time.Time `gorm:"timestamp"`          // 时间戳
	SeqID     int64     `gorm:"column:seq_id"`      // 用于保序的序列号
	gorm.Model
}

func (Conversation) TableName() string {
	return "conversation"
}
