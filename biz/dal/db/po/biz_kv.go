package po

import "github.com/jinzhu/gorm"

// BizKV 业务用的数据字典
type BizKV struct {
	ID      int64  `gorm:"column:id" json:"id"`           // 主键自增id, 无业务意义
	Key     string `gorm:"column:key" json:"key"`         // 键
	Val     string `gorm:"column:val" json:"val"`         // 值
	Creator int64  `gorm:"column:creator" json:"creator"` // 创建者的user_id
	Updater int64  `gorm:"column:updater" json:"updater"` // 更新者的user_id
	Deleter int64  `gorm:"column:deleter" json:"deleter"` // 删除者的user_id
	gorm.Model
}

func (BizKV) TableName() string {
	return "biz_kv"
}
