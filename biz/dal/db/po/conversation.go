package po

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 会话表
type Conversation struct {
	ID        int64     `gorm:"column:id"`        // 主键自增id,无业务意义
	ConvID    int64     `gorm:"column:conv_id"`   // 会话id
	Type      string    `gorm:"column:type"`      // 会话类型
	Creator   int64     `gorm:"column:creator"`   // 创建人id
	Status    string    `gorm:"column:status"`    // 会话状态
	Timestamp time.Time `gorm:"column:timestamp"` // 会话时间戳
	gorm.Model
}

func (Conversation) TableName() string {
	return "conversation"
}
