package po

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// UserConvRel 用户会话关系表
type UserConvRel struct {
	ID           int64         `gorm:"column:id"`           // 主键自增id,无业务意义
	RelID        int64         `gorm:"column:rel_id"`       // 关系id
	ConvID       int64         `gorm:"column:conv_id"`      // 会话id
	UserID       int64         `gorm:"column:user_id"`      // 用户id
	Participants pq.Int64Array `gorm:"column:participants"` // 会话参与者
	Role         string        `gorm:"column:role"`         // 用户角色
	gorm.Model
}

func (UserConvRel) TableName() string {
	return "user_conv_rel"
}
