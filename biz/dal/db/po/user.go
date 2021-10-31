package po

import "github.com/jinzhu/gorm"

/*
	用户表
*/
type User struct {
	ID       int64  `gorm:"column:id"`       // 主键自增id,无业务意义
	UserID   int64  `gorm:"column:user_id"`  // 用户id, 用户唯一标示
	Nickname string `gorm:"column:nickname"` // 用户昵称
	HeadURL  string `gorm:"column:head_url"` // 用户头像链接
	gorm.Model
}

func (User) TableName() string {
	return "user_center"
}

/*
	用户信息表
*/
type UserAccount struct {
	ID          int64  `gorm:"column:id"`           // 主键自增id,无业务意义
	UserID      int64  `gorm:"column:user_id"`      // 用户id, 用户唯一标示
	UserName    string `gorm:"column:username"`     // 用户登陆名
	PhoneNumber string `gorm:"column:phone_number"` // 用户手机号
	Email       string `gorm:"column:email"`        // 用户邮箱
	WxOpenID    string `gorm:"column:wx_open_id"`   // 微信open_id
	gorm.Model
}

func (UserAccount) TableName() string {
	return "user_account"
}
