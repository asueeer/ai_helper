package po

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Moment 动态表
type Moment struct {
	ID           int64           `json:"id"`
	MomentID     int64           `json:"moment_id"`     // 动态id
	UserID       int64           `json:"user_id"`       // 用户id
	Content      json.RawMessage `json:"content"`       // 动态内容
	ViewCount    int64           `json:"view_count"`    // 浏览量
	CommentCount int64           `json:"comment_count"` // 评论数
	LikeCount    int64           `json:"like_count"`    // 点赞数

	gorm.Model
}

func (Moment) TableName() string {
	return "moment"
}
