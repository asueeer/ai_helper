package po

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Comment 评论表
type Comment struct {
	ID           int64           `gorm:"column:id" json:"id"`                       // 主键自增id,无业务意义
	CommentID    int64           `gorm:"column:comment_id" json:"comment_id"`       // 评论id
	UserID       int64           `gorm:"column:user_id" json:"user_id"`             // (创建者)用户id
	EntityID     int64           `gorm:"column:entity_id" json:"entity_id"`         // 被评论的实体id
	Content      json.RawMessage `gorm:"column:content" json:"content"`             // 评论的内容(json格式)
	EntityType   string          `gorm:"column:entity_type" json:"entity_type"`     // 被评论的实体类型, 枚举值, "评论": comment; "动态": moment
	LikeCount    int64           `gorm:"column:like_count" json:"like_count"`       // 点赞数量
	CommentCount int64           `gorm:"column:comment_count" json:"comment_count"` // 评论数量

	gorm.Model
}

func (Comment) TableName() string {
	return "comment"
}
